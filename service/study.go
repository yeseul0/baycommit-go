package service

import (
	"baycommit-go/types"
	"time"
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
