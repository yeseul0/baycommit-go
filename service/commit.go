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

	// 이 레포가 등록된 스터디들 조회
	studies, err := GetStudiesByRepoUrl(commit.RepositoryUrl)
	if err != nil {
		return fmt.Errorf("스터디 조회 실패: %v", err)
	}
	if len(studies) == 0 {
		fmt.Printf("해당 레포에 등록된 스터디 없음")
		return nil
	}
	fmt.Printf("등록된 스터디 %d개 발견 \n", len(studies))

	//각 스터디마다 처리! (goroutine)
	for _, study := range studies {
		//<1> 스터디 날짜 계산(studyDate (UNIX) 게산)
		studyDate := CalculateStudyDate(
			commitTime,
			uint64(study.StudyStartTime),
			uint64(study.StudyEndTime),
		)
		if studyDate == 0 {
			fmt.Printf("스터디 %s: 스터디 시간 아님, skip\n", study.ProxyAddress)
			continue
		}
		fmt.Printf("스터디 %s: studyDate(UNIX) = %d\n", study.ProxyAddress, studyDate)

		//<2> 유저 지갑 주소 조회 (이 스터디에서 이 author는 무슨 지갑?)
		walletAddr, err := GetWalletAddress(commit.Author.Email, study.ProxyAddress)
		if err != nil {
			fmt.Printf("지갑 주소 조회 실패 : %v\n", err)
			continue
		}
		fmt.Printf("유저 지갑: %s\n", walletAddr)

		//블록체인 처리
		//<3> blockchain.GetStudyDayInfo() 호출 → 오늘 스터디 시작됐나?
		participantCount, _, err := GetStudyDayInfo(study.ProxyAddress, studyDate)
		if err != nil {
			fmt.Printf("스터디 정보 조회 실패: %v\n", err)
			continue
		}
		if participantCount == 0 { //스터디 시작 안됐으면 시작 (blockchain.StartTodayStudy() 호출)
			fmt.Println("오늘 스터디 시작 중...")
			err = StartTodayStudy(study.ProxyAddress, studyDate)
			if err != nil {
				fmt.Printf("스터디 시작 실패: %v\n", err)
				continue
			}
		}

		//<4> blockchain.GetCommitTime() 호출 → 이 유저 오늘 커밋했나?
		existingCommit, err := GetCommitTime(study.ProxyAddress, studyDate, walletAddr)
		if err != nil {
			fmt.Printf("커밋 시간 조회 실패: %v\n", err)
			continue
		}
		if existingCommit != 0 {
			fmt.Printf("이미 커밋 기록됨, skip\n")
			continue
		}
		fmt.Println("커밋 기록 중...")
		err = TrackCommit(study.ProxyAddress, studyDate, walletAddr, commitTime)
		if err != nil {
			fmt.Printf("커밋 기록 실패: %v\n", err)
			continue
		}

		fmt.Printf("커밋 기록 완료! 스터디 : %s\n", study.ProxyAddress)
		fmt.Printf("studyDate : %d\nwalletAddress : %s\n", studyDate, walletAddr)
	}

	// 1. CalculateStudyDate() 호출 → studyDate 계산

	// 2. blockchain.GetStudyDayInfo() 호출 → 오늘 스터디 시작됐나?
	// 3. blockchain.GetCommitTime() 호출 → 이 유저 오늘 커밋했나?
	// 4. blockchain.StartTodayStudy() 호출 (필요시)
	// 5. blockchain.TrackCommit() 호출 (필요시)
	return nil
}
