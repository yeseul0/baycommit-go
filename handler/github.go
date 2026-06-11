package handler

import (
	"baycommit-go/service"
	"baycommit-go/types"
	"encoding/json"
	"fmt"
	"net/http"
)

/*
GET /github/repositories
JWT 쿠키로 유저 식별 → DB에서 github_access_token 가져와 GitHub API 호출
*/
func GithubRepositoriesHandler(w http.ResponseWriter, r *http.Request) {
	// ① JWT에서 이메일 추출
	email, err := extractEmailFromCookie(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	// ② 이메일로 유저 조회 (github_access_token 포함)
	user, err := service.GetUserByEmail(email)
	if err != nil {
		http.Error(w, "유저를 찾을 수 없습니다", http.StatusUnauthorized)
		return
	}

	if user.GithubAccessToken == "" {
		http.Error(w, "GitHub 토큰이 없습니다. 다시 로그인하세요.", http.StatusUnauthorized)
		return
	}

	// ③ GitHub API로 레포 목록 조회
	repos, err := service.GetUserRepositories(user.GithubAccessToken)
	if err != nil {
		fmt.Printf("GitHub 레포 조회 실패: %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(types.GitHubRepositoriesResponse{
		Success:      true,
		Repositories: repos,
	})
}
