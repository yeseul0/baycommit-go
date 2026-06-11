package service

import (
	"baycommit-go/contract"
	"baycommit-go/types"
	"context"
	"fmt"
	"math/big"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

func GetStudyList(currentUserID uint) ([]types.StudyListItem, error) {
	// 전체 스터디 조회
	var studies []types.Study
	if err := DB.Find(&studies).Error; err != nil {
		return nil, err
	}

	today := time.Now().UTC().Format("2006-01-02")
	result := make([]types.StudyListItem, 0, len(studies))

	for _, study := range studies {
		item := types.StudyListItem{
			ID:             study.ID,
			StudyName:      study.Name,
			CreatedAt:      time.Unix(study.CreatedAt, 0).UTC().Format(time.RFC3339),
			ProxyAddress:   study.ProxyAddress,
			StudyStartTime: study.StudyStartTime,
			StudyEndTime:   study.StudyEndTime,
			IsOwner:        study.OwnerID == currentUserID,
		}

		// 내 참여 정보 조회
		var myUserStudy types.UserStudy
		err := DB.Where("study_id = ? AND user_id = ?", study.ID, currentUserID).
			First(&myUserStudy).Error
		if err == nil {
			item.IsParticipating = true
			item.MyRepoUrl = myUserStudy.RepoUrl
			item.HasRegisteredRepository = myUserStudy.RepoUrl != ""
		}

		// 참여자 수
		var count int64
		DB.Model(&types.UserStudy{}).Where("study_id = ?", study.ID).Count(&count)
		item.ParticipantCount = int(count)

		// 오늘 세션 조회
		var session types.StudySession
		err = DB.Where("study_id = ? AND study_date = ?", study.ID, today).
			First(&session).Error
		if err == nil {
			todaySession := &types.TodaySession{
				SessionId: session.ID,
				Status:    string(session.Status),
				StudyDate: session.StudyDate,
				StartedAt: session.StartedAt,
			}

			// 참여자 목록 조회
			var userStudies []types.UserStudy
			DB.Where("study_id = ?", study.ID).Find(&userStudies)

			for _, us := range userStudies {
				var user types.User
				if err := DB.First(&user, us.UserID).Error; err != nil {
					continue
				}

				participant := types.ParticipantInfo{
					GithubEmail:   user.GithubEmail,
					GithubHandle:  user.GithubLogin,
					WalletAddress: user.WalletAddress,
				}

				// 커밋 기록 조회
				var commitRecord types.CommitRecord
				err := DB.Where("study_session_id = ? AND user_id = ?", session.ID, us.UserID).
					First(&commitRecord).Error
				if err == nil {
					participant.Commit = &types.CommitInfo{
						CommitId:      commitRecord.CommitID,
						CommitMessage: commitRecord.CommitMessage,
						CommitTime:    commitRecord.CommitTimestamp,
					}
				}

				todaySession.Participants = append(todaySession.Participants, participant)
			}

			item.TodaySession = todaySession
		}

		result = append(result, item)
	}

	return result, nil
}

/*
GET /study/all/commits/today
오늘 세션이 있는 스터디만, 참여자 + 커밋 현황 반환
*/
func GetTodayCommits() ([]types.TodayCommitStudy, error) {
	var studies []types.Study
	if err := DB.Find(&studies).Error; err != nil {
		return nil, err
	}

	today := time.Now().UTC().Format("2006-01-02")
	result := make([]types.TodayCommitStudy, 0)

	for _, study := range studies {
		// 오늘 세션 조회
		var sessions []types.StudySession
		DB.Where("study_id = ? AND study_date = ?", study.ID, today).Find(&sessions)
		if len(sessions) == 0 {
			continue // 오늘 세션 없는 스터디는 제외
		}

		studyItem := types.TodayCommitStudy{
			StudyName:    study.Name,
			ProxyAddress: study.ProxyAddress,
		}

		for _, session := range sessions {
			sessionItem := types.TodayCommitSession{
				SessionId: session.ID,
				StudyDate: session.StudyDate,
				Status:    string(session.Status),
				StartedAt: session.StartedAt,
			}

			// 참여자 목록
			var userStudies []types.UserStudy
			DB.Where("study_id = ?", study.ID).Find(&userStudies)

			for _, us := range userStudies {
				var user types.User
				if err := DB.First(&user, us.UserID).Error; err != nil {
					continue
				}

				participant := types.ParticipantInfo{
					GithubEmail:   user.GithubEmail,
					GithubHandle:  user.GithubLogin,
					WalletAddress: user.WalletAddress,
				}

				// 커밋 기록
				var commitRecord types.CommitRecord
				err := DB.Where("study_session_id = ? AND user_id = ?", session.ID, us.UserID).
					First(&commitRecord).Error
				if err == nil {
					participant.Commit = &types.CommitInfo{
						CommitId:      commitRecord.CommitID,
						CommitMessage: commitRecord.CommitMessage,
						CommitTime:    commitRecord.CommitTimestamp,
					}
				}

				sessionItem.Participants = append(sessionItem.Participants, participant)
			}

			studyItem.Sessions = append(studyItem.Sessions, sessionItem)
		}

		result = append(result, studyItem)
	}

	return result, nil
}

/*
POST /study/create
Factory 컨트랙트로 스터디 생성 → proxyAddress 받아 DB 저장
*/
func CreateStudy(ownerID uint, req types.StudyCreateRequest) (types.Study, string, error) {
	// ① Factory 컨트랙트 인스턴스
	factoryAddr := common.HexToAddress(os.Getenv("FACTORY_ADDRESS"))
	factory, err := contract.NewFactory(factoryAddr, client)
	if err != nil {
		return types.Study{}, "", fmt.Errorf("Factory 인스턴스 생성 실패: %v", err)
	}

	// ② 금액 파싱 (문자열 → *big.Int)
	depositAmount := new(big.Int)
	if _, ok := depositAmount.SetString(req.DepositAmount, 10); !ok {
		return types.Study{}, "", fmt.Errorf("depositAmount 파싱 실패: %s", req.DepositAmount)
	}
	penaltyAmount := new(big.Int)
	if _, ok := penaltyAmount.SetString(req.PenaltyAmount, 10); !ok {
		return types.Study{}, "", fmt.Errorf("penaltyAmount 파싱 실패: %s", req.PenaltyAmount)
	}

	// ③ studyAdmin: 서버 admin 주소 사용 (onlyAdmin이 createProxy 호출하므로)
	adminAddr := common.HexToAddress(os.Getenv("STUDY_ADMIN_ADDRESS"))

	// ④ 트랜잭션 옵션
	auth, err := getTxOpts()
	if err != nil {
		return types.Study{}, "", fmt.Errorf("트랜잭션 옵션 생성 실패: %v", err)
	}

	// ⑤ createProxy 호출
	tx, err := factory.CreateProxy(
		auth,
		req.StudyName,
		depositAmount,
		penaltyAmount,
		adminAddr,
		big.NewInt(req.StudyStartTime),
		big.NewInt(req.StudyEndTime),
	)
	if err != nil {
		return types.Study{}, "", fmt.Errorf("createProxy 트랜잭션 실패: %v", err)
	}
	fmt.Printf("createProxy tx 전송됨: %s\n", tx.Hash().Hex())

	// ⑥ 트랜잭션 채굴 대기
	receipt, err := bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		return types.Study{}, "", fmt.Errorf("트랜잭션 마이닝 대기 실패: %v", err)
	}
	if receipt.Status == 0 {
		return types.Study{}, "", fmt.Errorf("트랜잭션이 블록체인에서 Revert됨 (tx: %s)", tx.Hash().Hex())
	}

	// ⑦ 이벤트 로그에서 proxyAddress 추출
	proxyAddress, err := parseProxyCreatedEvent(factory, receipt.Logs)
	if err != nil {
		return types.Study{}, "", fmt.Errorf("ProxyCreated 이벤트 파싱 실패: %v", err)
	}
	fmt.Printf("생성된 Proxy 주소: %s\n", proxyAddress)

	// ⑧ DB에 스터디 저장
	study := types.Study{
		Name:           req.StudyName,
		OwnerID:        ownerID,
		ProxyAddress:   proxyAddress,
		StudyStartTime: req.StudyStartTime,
		StudyEndTime:   req.StudyEndTime,
	}
	if err := DB.Create(&study).Error; err != nil {
		return types.Study{}, "", fmt.Errorf("스터디 DB 저장 실패: %v", err)
	}

	// ⑨ 오너를 user_studies에 참여자로 등록
	userStudy := types.UserStudy{
		UserID:  ownerID,
		StudyID: study.ID,
	}
	if err := DB.Create(&userStudy).Error; err != nil {
		return types.Study{}, "", fmt.Errorf("user_studies 저장 실패: %v", err)
	}

	return study, tx.Hash().Hex(), nil
}

// receipt 로그에서 ProxyCreated 이벤트를 파싱해 proxy 주소 반환
func parseProxyCreatedEvent(factory *contract.Factory, logs []*ethtypes.Log) (string, error) {
	for _, log := range logs {
		event, err := factory.FactoryFilterer.ParseProxyCreated(*log)
		if err != nil {
			continue // 다른 컨트랙트 이벤트는 skip
		}
		return event.ProxyAddress.Hex(), nil
	}
	return "", fmt.Errorf("ProxyCreated 이벤트를 찾지 못함")
}
