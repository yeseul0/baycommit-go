package service

import (
	"baycommit-go/types"
	"fmt"
	"time"
)

// ISO 8601 문자열 → Unix timestamp
func ParseCommitTime(timestamp string) (uint64, error) {
	t, err := time.Parse(time.RFC3339, timestamp)
	if err != nil {
		return 0, err
	}
	return uint64(t.Unix()), nil
}

func ProcessCommit(commit types.Commit) error {
	//0. 커밋 시간 파싱
	commitTime, err := ParseCommitTime(commit.Timestamp) //commitTime unix time임
	if err != nil {
		return fmt.Errorf("커밋 시간 파싱 실패: %v", err)
	}
	fmt.Printf("커밋 시간: %d\n", commitTime)

	// repo_url로 스터디 + 유저 동시 특정
	userStudy, study, err := GetUserStudyByRepoUrl(commit.RepositoryUrl)
	if err != nil {
		fmt.Printf("해당 레포에 등록된 스터디 없음")
		return nil
	}
	fmt.Printf("스터디 발견: %s\n", study.ProxyAddress)

	//<1> 스터디 날짜 계산
	studyDate := CalculateStudyDate(
		commitTime,
		uint64(study.StudyStartTime),
		uint64(study.StudyEndTime),
	)
	if studyDate == 0 {
		fmt.Printf("스터디 %s: 스터디 시간 아님, skip\n", study.ProxyAddress)
		return nil
	}
	fmt.Printf("스터디 %s: studyDate(UNIX) = %d\n", study.ProxyAddress, studyDate)

	//<2> 유저 지갑 주소 조회
	walletAddr, err := GetWalletAddressByUserID(userStudy.UserID)
	if err != nil {
		return fmt.Errorf("지갑 주소 조회 실패: %v", err)
	}
	fmt.Printf("유저 지갑: %s\n", walletAddr)

	//블록체인 처리
	//<3> 오늘 스터디 시작됐나?
	participantCount, _, err := GetStudyDayInfo(study.ProxyAddress, studyDate)
	if err != nil {
		return fmt.Errorf("스터디 정보 조회 실패: %v", err)
	}
	if participantCount == 0 {
		fmt.Println("오늘 스터디 시작 중...")
		if err := StartTodayStudy(study.ProxyAddress, studyDate); err != nil {
			return fmt.Errorf("스터디 시작 실패: %v", err)
		}
	}

	//<4> 이 유저 오늘 커밋했나?
	existingCommit, err := GetCommitTime(study.ProxyAddress, studyDate, walletAddr)
	if err != nil {
		return fmt.Errorf("커밋 시간 조회 실패: %v", err)
	}
	if existingCommit != 0 {
		fmt.Printf("이미 커밋 기록됨, skip\n")
		// [메트릭] 커밋 처리 횟수 - skip
		CommitProcessTotal.WithLabelValues("skip").Inc()
		return nil
	}

	fmt.Println("커밋 기록 중...")
	if err := TrackCommit(study.ProxyAddress, studyDate, walletAddr, commitTime); err != nil {
		// [메트릭] 커밋 처리 횟수 - error
		CommitProcessTotal.WithLabelValues("error").Inc()
		return fmt.Errorf("커밋 기록 실패: %v", err)
	}

	// [메트릭] 커밋 처리 횟수 - success
	CommitProcessTotal.WithLabelValues("success").Inc()
	fmt.Printf("커밋 기록 완료! 스터디: %s\n", study.ProxyAddress)
	fmt.Printf("studyDate: %d\nwalletAddress: %s\n", studyDate, walletAddr)
	return nil
}
