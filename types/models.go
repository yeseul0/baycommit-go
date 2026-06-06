package types

// ===== DB Models (GORM) =====
// studies 테이블
type Study struct {
	ID             uint   `gorm:"primaryKey"`
	ProxyAddress   string `gorm:"column:proxy_address"`
	StudyStartTime int64  `gorm:"column:study_start_time"`
	StudyEndTime   int64  `gorm:"column:study_end_time`
}

func (Study) TableName() string { //Study 구조체는 studies 테이블임
	return "studies"
}

// users 테이블
type User struct {
	ID            uint   `gorm:"primaryKey"`
	GithubEmail   string `gorm:"column:github_email"`
	WalletAddress string `gorm:"column:wallet_address"`
}

func (User) TableName() string {
	return "users"
}

// user_studies 테이블
type UserStudy struct {
	ID      uint   `gorm:"primaryKey"`
	UserID  uint   `gorm:"column:user_id"`
	StudyID uint   `gorm:"column:study_id"`
	RepoUrl string `gorm:"column:repo_url"`
}

func (UserStudy) TableName() string {
	return "user_studies"
}

// study_sessions 테이블
type StudySession struct {
	ID               uint          `gorm:"primaryKey"`
	StudyID          uint          `gorm:"column:study_id"`
	StudyDate        string        `gorm:"column:study_date"` //YYYY-MM-DD 형식
	Status           SessionStatus `gorm:"column:status"`     // PENDING, ACTIVE, CLOSED, FAILED
	StartedAt        int64         `gorm:"column:started_at"`
	ClosedAt         int64         `gorm:"column:closed_at"`
	BlockchainTxHash string        `gorm:"column:blockchain_tx_hash"`
	StudyMidnightUtc int64         `gorm:"column:study_midnight_utc"` //스터디 날짜 자정 (UTC)
}

func (StudySession) TableName() string {
	return "study_sessions"
}

// commit_records 테이블
type CommitRecord struct {
	ID              uint   `gorm:"primaryKey"`
	StudySessionID  uint   `gorm:"column:study_session_id"`
	UserID          uint   `gorm:"column:user_id"`
	CommitTimestamp int64  `gorm:"column:commit_timestamp"`
	CommitID        string `gorm:"column:commit_id"`
	CommitMessage   string `gorm:"column:commit_message"`
}

func (CommitRecord) TableName() string {
	return "commit_records"
}
