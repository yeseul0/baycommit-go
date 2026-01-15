package service

import "baycommit-go/types"

func ProcessCommit(commit types.Commit) error {
	// 1. CalculateStudyDate() 호출 → studyDate 계산
	// 2. blockchain.GetStudyDayInfo() 호출 → 오늘 스터디 시작됐나?
	// 3. blockchain.GetCommitTime() 호출 → 이 유저 오늘 커밋했나?
	// 4. blockchain.StartTodayStudy() 호출 (필요시)
	// 5. blockchain.TrackCommit() 호출 (필요시)
	return nil
}
