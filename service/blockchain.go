package service

// 컨트랙트 조회/호출?

// 컨트랙트 조회: 오늘 스터디 정보
func GetStudyDayInfo(proxyAddr string, timestamp uint64) (participantCount uint64, isClosed bool, err error) {
	//
	return 0, false, nil
}

// 컨트랙트 조회: 유저 커밋 시간
func GetCommitTime(proxyAddr string, timestamp uint64, userAddr string) (uint64, error) {
	//
	return 0, nil
}

// 컨트랙트 호출 : 오늘 스터디 시작
func StartTodayStudy(proxyAddr string, timestamp uint64) error {
	//
	return nil
}

// 컨트랙트 호출 : 커밋 기록
func TrackCommit(proxyAddr string, timestamp uint64, userAddr string, commitTime uint64) error {
	//
	return nil
}
