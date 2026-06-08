package handler

import (
	"baycommit-go/service"
	"baycommit-go/types"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

/*
GET /study/list
쿠키의 access_token JWT로 현재 유저 식별 후 스터디 목록 반환
*/
func StudyListHandler(w http.ResponseWriter, r *http.Request) {
	// JWT에서 이메일 추출
	cookie, err := r.Cookie("access_token")
	if err != nil {
		http.Error(w, "로그인이 필요합니다", http.StatusUnauthorized)
		return
	}

	token, err := jwt.Parse(cookie.Value, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil || !token.Valid {
		http.Error(w, "유효하지 않은 토큰", http.StatusUnauthorized)
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		http.Error(w, "토큰 파싱 실패", http.StatusUnauthorized)
		return
	}

	email, ok := claims["email"].(string)
	if !ok || email == "" {
		http.Error(w, "토큰에 이메일 없음", http.StatusUnauthorized)
		return
	}

	// 이메일로 유저 조회
	currentUser, err := service.GetUserByEmail(email)
	if err != nil {
		fmt.Printf("유저 조회 실패: %v\n", err)
		http.Error(w, "유저를 찾을 수 없습니다", http.StatusUnauthorized)
		return
	}

	// 스터디 목록 조회
	studies, err := service.GetStudyList(currentUser.ID)
	if err != nil {
		fmt.Printf("스터디 목록 조회 실패: %v\n", err)
		http.Error(w, "스터디 목록 조회 실패", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(types.StudyListResponse{
		Success: true,
		Studies: studies,
	})
}
