package service

import (
	"baycommit-go/types"
	"os"

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
