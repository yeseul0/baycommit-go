package service

import (
	"baycommit-go/types"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

/*
① code → GitHub access token 교환
*/
func ExchangeCodeForToken(code string) (string, error) {
	body, _ := json.Marshal(map[string]string{
		"client_id":     os.Getenv("GITHUB_CLIENT_ID"),
		"client_secret": os.Getenv("GITHUB_CLIENT_SECRET"),
		"code":          code,
	})

	req, _ := http.NewRequest("POST", "https://github.com/login/oauth/access_token", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("GitHub 토큰 요청 실패: %v", err)
	}
	defer resp.Body.Close()

	var tokenResp types.GitHubTokenResponse
	json.NewDecoder(resp.Body).Decode(&tokenResp)

	if tokenResp.AccessToken == "" {
		return "", fmt.Errorf("access_token이 비어있음")
	}
	return tokenResp.AccessToken, nil
}

/*
② access token → GitHub 유저 정보 조회
이메일이 비공개인 경우 /user/emails 로 fallback
*/
func GetGitHubUser(accessToken string) (types.GitHubUser, error) {
	req, _ := http.NewRequest("GET", "https://api.github.com/user", nil)
	req.Header.Set("Authorization", "Bearer "+accessToken)
	req.Header.Set("Accept", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return types.GitHubUser{}, fmt.Errorf("GitHub 유저 정보 요청 실패: %v", err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	var user types.GitHubUser
	json.Unmarshal(body, &user)

	// 이메일 비공개인 경우 /user/emails에서 primary 이메일 조회
	if user.Email == "" {
		email, err := getPrimaryEmail(accessToken)
		if err != nil {
			return user, fmt.Errorf("GitHub 이메일 조회 실패: %v", err)
		}
		user.Email = email
	}

	return user, nil
}

func getPrimaryEmail(accessToken string) (string, error) {
	req, _ := http.NewRequest("GET", "https://api.github.com/user/emails", nil)
	req.Header.Set("Authorization", "Bearer "+accessToken)
	req.Header.Set("Accept", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var emails []types.GitHubEmail
	json.NewDecoder(resp.Body).Decode(&emails)

	for _, e := range emails {
		if e.Primary && e.Verified {
			return e.Email, nil
		}
	}
	return "", fmt.Errorf("primary 이메일 없음")
}

/*
③ DB에 유저 저장 (없으면 create, 있으면 access token 업데이트)
*/
func UpsertUser(githubEmail string, githubLogin string, accessToken string) error {
	var user types.User
	result := DB.Where("github_email = ?", githubEmail).First(&user)
	if result.Error != nil {
		// 신규 유저 생성
		user = types.User{
			GithubEmail:       githubEmail,
			GithubLogin:       githubLogin,
			GithubAccessToken: accessToken,
		}
		return DB.Create(&user).Error
	}
	// 기존 유저 → access token 갱신 (재로그인 시 토큰 바뀔 수 있음)
	return DB.Model(&user).Updates(map[string]interface{}{
		"github_login":        githubLogin,
		"github_access_token": accessToken,
	}).Error
}

/*
④ JWT 발급
*/
func IssueJWT(githubEmail string) (string, error) {
	secret := os.Getenv("JWT_SECRET")
	claims := jwt.MapClaims{
		"email": githubEmail,
		"exp":   time.Now().Add(7 * 24 * time.Hour).Unix(), // 7일
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}
