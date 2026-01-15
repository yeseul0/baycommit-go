package types

// ===== DB Models (GORM) =====
type Study struct {
	ID             uint   `gorm:"primaryKey"`
	ProxyAddress   string `gorm:"column:proxy_address"`
	StudyStartTime int64  `gorm:"column:study_start_time"`
	StudyEndTime   int64  `gorm:"column:study_end_time`
}

func (Study) TableName() string { //Study 구조체는 studies 테이블임
	return "studies"
}

type User struct {
	ID          uint   `gorm:"primaryKey"`
	GithubEmail string `gorm:"column:github_email"`
}

func (User) TableName() string {
	return "users"
}

type UserStudy struct {
	ID            uint   `gorm:"primaryKey"`
	UserID        uint   `gorm:"column:user_id"`
	StudyID       uint   `gorm:"column:study_id"`
	WalletAddress string `gorm:"columnc:wallet_address"`
}

func (UserStudy) TableName() string {
	return "user_studies"
}

type Repository struct {
	ID       uint `gorm:"primaryKey"`
	StudyID  uint `gorm:"column:study_id"`
	RepoUrl  uint `gorm:"column:repo_url"`
	IsActive bool `gorm:"column:is_active"`
}

func (Repository) TableName() string {
	return "repositories"
}
