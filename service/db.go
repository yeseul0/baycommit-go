package service

import (
	"baycommit-go/types"
	"fmt"
	"os"

	"time"

	"gorm.io/driver/postgres" //PostgreSQL 드라이버
	"gorm.io/gorm"
)

var DB *gorm.DB

// DB 연결
func ConnectDB() error {
	dsn := os.Getenv("DATABASE_URL") //Data Src Name
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	return err
}

// repourl -> study 컨트랙트s (이 레포가 등록된 스터디들)
func GetStudiesByRepoUrl(repoUrl string) ([]types.Study, error) {
	var studies []types.Study

	err := DB.
		Joins("JOIN repositories ON repositories.study_id = studies.id").
		Where("repositories.repo_url = ? AND repositories.is_active = true", repoUrl). //SQL injection 방지
		Find(&studies).Error

	return studies, err
}

// 이메일 + 프록시 주소 -> 지갑주소 (왜냐면 한 유저가 지갑 다르게 여러 스터디에 참여할수도 있음)
func GetWalletAddress(email, proxyAddress string) (string, error) { //하나만 찾는 함수임! 배열 아님
	var userStudy types.UserStudy
	err := DB.
		Joins("JOIN users ON users.id = user_studies.user_id").
		Joins("JOIN studies ON studies.id = user_studies.study_id").
		Where("users.github_email = ? AND studies.proxy_address = ?", email, proxyAddress).
		First(&userStudy).Error

	if err != nil {
		return "", err
	}
	return userStudy.WalletAddress, nil
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
}
