package service

import (
	"baycommit-go/types"
	"fmt"
	"os"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"gorm.io/driver/postgres" //PostgreSQL 드라이버
	"gorm.io/gorm"
)

var DB *gorm.DB

// DB 연결
func ConnectDB() error {
	dsn := os.Getenv("DATABASE_URL") //Data Src Name
	var err error
	DB, err = gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true, // PgBouncer(Supabase) 호환
	}), &gorm.Config{})
	return err
}

// email -> 유저 조회
func GetUserByEmail(email string) (types.User, error) {
	// [메트릭] DB 쿼리 소요시간 측정 - get_user_by_email
	timer := prometheus.NewTimer(DBQueryDuration.WithLabelValues("get_user_by_email"))
	defer timer.ObserveDuration()

	var user types.User
	err := DB.Where("github_email = ?", email).First(&user).Error
	return user, err
}

// repo_url -> (study, user) 동시 특정
func GetUserStudyByRepoUrl(repoUrl string) (types.UserStudy, types.Study, error) {
	// [메트릭] DB 쿼리 소요시간 측정 - get_user_study
	timer := prometheus.NewTimer(DBQueryDuration.WithLabelValues("get_user_study"))
	defer timer.ObserveDuration()

	var userStudy types.UserStudy
	var study types.Study

	err := DB.
		Where("repo_url = ?", repoUrl).
		First(&userStudy).Error
	if err != nil {
		return userStudy, study, err
	}

	err = DB.First(&study, userStudy.StudyID).Error
	return userStudy, study, err
}

// user_id -> 지갑주소
func GetWalletAddressByUserID(userID uint) (string, error) {
	// [메트릭] DB 쿼리 소요시간 측정 - get_wallet
	timer := prometheus.NewTimer(DBQueryDuration.WithLabelValues("get_wallet"))
	defer timer.ObserveDuration()

	var user types.User
	err := DB.First(&user, userID).Error
	if err != nil {
		return "", err
	}
	return user.WalletAddress, nil
}

// 새로운 세션 PENDING 상태로 미리 생성
func CreatePendingSession(proxyAddr string, studyDateUnix uint64, txHash string) error {
	// 1. ProxyAddress로 Study ID 찾기
	var study types.Study
	err := DB.Where("proxy_address = ?", proxyAddr).First(&study).Error
	if err != nil {
		return fmt.Errorf("스터디 찾기 실패: %v", err)
	}

	t := time.Unix(int64(studyDateUnix), 0).UTC()
	dateString := t.Format("2006-01-02")

	// 2. SELECT 먼저!! 검색 조건 (중복 방지)
	where := types.StudySession{
		StudyID:          study.ID,
		StudyMidnightUtc: int64(studyDateUnix),
	}

	// 3. 생성 데이터 (PENDING 상태)
	attrs := types.StudySession{
		StudyDate:        dateString,
		Status:           types.StatusPending,
		StartedAt:        time.Now().Unix(),
		BlockchainTxHash: txHash,
	}

	// 4. 저장 (FirstOrCreate)
	return DB.Where(where).Attrs(attrs).FirstOrCreate(&types.StudySession{}).Error
}

// 스터디 참여 (user_studies INSERT + wallet_address 업데이트)
func JoinStudy(userID uint, proxyAddr string, walletAddress string) error {
	// study 조회
	var study types.Study
	if err := DB.Where("proxy_address = ?", proxyAddr).First(&study).Error; err != nil {
		return fmt.Errorf("스터디를 찾을 수 없습니다")
	}

	// 이미 참여 중인지 확인
	var existing types.UserStudy
	result := DB.Where("user_id = ? AND study_id = ?", userID, study.ID).First(&existing)
	if result.Error == nil {
		return fmt.Errorf("이미 참여 중인 스터디입니다")
	}

	// user_studies INSERT
	if err := DB.Create(&types.UserStudy{
		UserID:  userID,
		StudyID: study.ID,
	}).Error; err != nil {
		return fmt.Errorf("참여 등록 실패: %v", err)
	}

	// wallet_address 업데이트
	if err := DB.Model(&types.User{}).
		Where("id = ?", userID).
		Update("wallet_address", walletAddress).Error; err != nil {
		return fmt.Errorf("지갑 주소 저장 실패: %v", err)
	}

	return nil
}

// proxyAddress → 스터디 참여자 + repoUrl 목록
func GetStudyParticipantRepos(proxyAddr string) ([]types.StudyParticipantRepo, error) {
	var study types.Study
	if err := DB.Where("proxy_address = ?", proxyAddr).First(&study).Error; err != nil {
		return nil, fmt.Errorf("스터디를 찾을 수 없습니다")
	}

	var userStudies []types.UserStudy
	if err := DB.Where("study_id = ?", study.ID).Find(&userStudies).Error; err != nil {
		return nil, err
	}

	result := make([]types.StudyParticipantRepo, 0, len(userStudies))
	for _, us := range userStudies {
		var user types.User
		if err := DB.First(&user, us.UserID).Error; err != nil {
			continue
		}
		result = append(result, types.StudyParticipantRepo{
			ParticipantEmail:  user.GithubEmail,
			ParticipantHandle: user.GithubLogin,
			ParticipantWallet: user.WalletAddress,
			RepoUrl:           us.RepoUrl,
		})
	}
	return result, nil
}

// user_studies.repo_url 업데이트
func RegisterRepository(userID uint, proxyAddr string, repoUrl string) error {
	// proxyAddress → study 조회
	var study types.Study
	if err := DB.Where("proxy_address = ?", proxyAddr).First(&study).Error; err != nil {
		return fmt.Errorf("스터디를 찾을 수 없습니다: %v", err)
	}

	// 해당 유저의 user_studies 레코드에 repo_url 저장
	result := DB.Model(&types.UserStudy{}).
		Where("user_id = ? AND study_id = ?", userID, study.ID).
		Update("repo_url", repoUrl)
	if result.Error != nil {
		return fmt.Errorf("레포 등록 실패: %v", result.Error)
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("해당 스터디 참여 기록이 없습니다")
	}
	return nil
}

// 세션 status 업데이트
func UpdateStudySessionStatus(proxyAddr string, studyDateUnix uint64, status types.SessionStatus) error {
	// 1. ProxyAddress로 Study ID 찾기
	var study types.Study
	if err := DB.Where("proxy_address = ?", proxyAddr).First(&study).Error; err != nil {
		return err
	}

	// 2. 해당 스터디 + 날짜의 세션을 찾아서 상태 업데이트
	return DB.Model(&types.StudySession{}).
		Where("study_id = ? AND study_midnight_utc = ?", study.ID, int64(studyDateUnix)).
		Update("status", status).Error
} // 세션의 status 업데이트 (성공, 실패)
