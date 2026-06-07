package handler

import (
	"baycommit-go/service"
	"fmt"
	"net/http"
)

/*
GitHub OAuth 콜백 핸들러
GET /auth/github/callback?code=XXX&state=https://frontend.com/dashboard
*/
func GithubCallbackHandler(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	state := r.URL.Query().Get("state") // 프론트 리다이렉트 URL

	if code == "" {
		http.Error(w, "code가 없습니다", http.StatusBadRequest)
		return
	}

	// ① code → access token
	accessToken, err := service.ExchangeCodeForToken(code)
	if err != nil {
		fmt.Printf("토큰 교환 실패: %v\n", err)
		http.Error(w, "GitHub 인증 실패", http.StatusInternalServerError)
		return
	}

	// ② access token → GitHub 유저 정보
	githubUser, err := service.GetGitHubUser(accessToken)
	if err != nil {
		fmt.Printf("유저 정보 조회 실패: %v\n", err)
		http.Error(w, "GitHub 유저 정보 조회 실패", http.StatusInternalServerError)
		return
	}

	// ③ DB upsert
	err = service.UpsertUser(githubUser.Email)
	if err != nil {
		fmt.Printf("DB 저장 실패: %v\n", err)
		http.Error(w, "유저 저장 실패", http.StatusInternalServerError)
		return
	}

	// ④ JWT 발급
	jwtToken, err := service.IssueJWT(githubUser.Email)
	if err != nil {
		fmt.Printf("JWT 발급 실패: %v\n", err)
		http.Error(w, "JWT 발급 실패", http.StatusInternalServerError)
		return
	}

	// ⑤ JWT 쿠키로 set
	http.SetCookie(w, &http.Cookie{
		Name:     "access_token",
		Value:    jwtToken,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
		Path:     "/",
		MaxAge:   7 * 24 * 60 * 60, // 7일
	})

	// ⑥ 프론트로 리다이렉트
	if state == "" {
		state = "/"
	}
	http.Redirect(w, r, state, http.StatusFound)
}
