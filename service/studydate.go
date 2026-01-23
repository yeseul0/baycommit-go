package service

func CalculateStudyDate(commitTimestamp uint64, startTime, endTime uint64) uint64 {
	//커밋 시간 → 어느 날짜 스터디인지 계산
	todayMidnight := (commitTimestamp / 86400) * 86400 //
	secondsInDay := commitTimestamp % 86400

	// 자정이 넘나드는 스터디?
	if endTime >= 86400 {
		endSeconds := endTime % 86400 //26시 -> 2시로 ㅇㅇ

		// 새벽 커밋 → 전날 스터디 날짜 자정(UNIX) return
		if secondsInDay < endSeconds {
			return todayMidnight - 86400
		}
	}
	// 저녁 커밋 → 당일 스터디 날짜(unix) return
	if secondsInDay >= startTime {
		return todayMidnight
	}

	//스터디 시간 아님
	return 0
}
