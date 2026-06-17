package handler

import (
	"baycommit-go/service"
	"baycommit-go/types"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/prometheus/client_golang/prometheus"
)

// JWT 쿠키에서 이메일 추출 (공통 로직)
func extractEmailFromCookie(r *http.Request) (string, error) {
	cookie, err := r.Cookie("access_token")
	if err != nil {
		return "", fmt.Errorf("쿠키 없음")
	}
	token, err := jwt.Parse(cookie.Value, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil || !token.Valid {
		return "", fmt.Errorf("유효하지 않은 토큰")
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", fmt.Errorf("토큰 파싱 실패")
	}
	email, ok := claims["email"].(string)
	if !ok || email == "" {
		return "", fmt.Errorf("토큰에 이메일 없음")
	}
	return email, nil
}

/*
GET /study/list
쿠키의 access_token JWT로 현재 유저 식별 후 스터디 목록 반환
*/
func StudyListHandler(w http.ResponseWriter, r *http.Request) {
	email, err := extractEmailFromCookie(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	currentUser, err := service.GetUserByEmail(email)
	if err != nil {
		fmt.Printf("유저 조회 실패: %v\n", err)
		http.Error(w, "유저를 찾을 수 없습니다", http.StatusUnauthorized)
		return
	}

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

/*
GET /study/:proxyAddress/repositories
스터디 참여자 전체 + 각자 등록된 repoUrl 반환
*/
func StudyRepositoriesHandler(w http.ResponseWriter, r *http.Request) {
	// URL: /study/0xABC.../repositories → proxyAddress 추출
	// path = ["", "study", "0xABC...", "repositories"]
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) != 4 || parts[3] != "repositories" {
		http.Error(w, "잘못된 경로", http.StatusBadRequest)
		return
	}
	proxyAddress := parts[2]

	participants, err := service.GetStudyParticipantRepos(proxyAddress)
	if err != nil {
		fmt.Printf("참여자 레포 조회 실패: %v\n", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(types.StudyRepositoriesResponse{
		Success:      true,
		Participants: participants,
	})
}

/*
POST /study/join
스터디 참여 + 지갑 주소 등록 (생성자도 이걸 타야 함)
*/
func StudyJoinHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	email, err := extractEmailFromCookie(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	currentUser, err := service.GetUserByEmail(email)
	if err != nil {
		http.Error(w, "유저를 찾을 수 없습니다", http.StatusUnauthorized)
		return
	}

	var req types.StudyJoinRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "요청 파싱 실패", http.StatusBadRequest)
		return
	}
	if req.ProxyAddress == "" || req.WalletAddress == "" {
		http.Error(w, "proxyAddress, walletAddress 필수", http.StatusBadRequest)
		return
	}

	if err := service.JoinStudy(currentUser.ID, req.ProxyAddress, req.WalletAddress); err != nil {
		fmt.Printf("스터디 참여 실패: %v\n", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(types.StudyJoinResponse{
		Success: true,
		Message: "스터디에 참여했습니다.",
	})
}

/*
POST /study/repository/register
스터디 참여자가 본인 레포 등록
*/
func StudyRepositoryRegisterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	email, err := extractEmailFromCookie(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	currentUser, err := service.GetUserByEmail(email)
	if err != nil {
		http.Error(w, "유저를 찾을 수 없습니다", http.StatusUnauthorized)
		return
	}

	var req types.RepositoryRegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "요청 파싱 실패", http.StatusBadRequest)
		return
	}
	if req.ProxyAddress == "" || req.RepoUrl == "" {
		http.Error(w, "proxyAddress, repoUrl 필수", http.StatusBadRequest)
		return
	}

	if err := service.RegisterRepository(currentUser.ID, req.ProxyAddress, req.RepoUrl); err != nil {
		fmt.Printf("레포 등록 실패: %v\n", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(types.RepositoryRegisterResponse{
		Success: true,
		Message: "레포지토리가 등록되었습니다.",
	})
}

/*
GET /study/all/commits/today
인증 불필요 - 대시보드 실시간 업데이트용
*/
func TodayCommitsHandler(w http.ResponseWriter, r *http.Request) {
	studies, err := service.GetTodayCommits()
	if err != nil {
		fmt.Printf("오늘 커밋 조회 실패: %v\n", err)
		http.Error(w, "오늘 커밋 조회 실패", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(types.TodayCommitsResponse{
		Success: true,
		Studies: studies,
	})
}

/*
POST /study/create
Factory 컨트랙트 호출로 스터디 생성
Body: { studyName, depositAmount, penaltyAmount, studyStartTime, studyEndTime }
*/
func StudyCreateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	// [메트릭] HTTP 전체 응답시간 측정 시작
	timer := prometheus.NewTimer(service.StudyCreateDuration)
	defer timer.ObserveDuration()

	// 인증
	email, err := extractEmailFromCookie(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	currentUser, err := service.GetUserByEmail(email)
	if err != nil {
		http.Error(w, "유저를 찾을 수 없습니다", http.StatusUnauthorized)
		return
	}

	// 요청 파싱
	var req types.StudyCreateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "요청 파싱 실패", http.StatusBadRequest)
		return
	}
	if req.StudyName == "" || req.DepositAmount == "" || req.PenaltyAmount == "" ||
		req.StudyStartTime == 0 || req.StudyEndTime == 0 {
		http.Error(w, "필수 필드 누락", http.StatusBadRequest)
		return
	}

	// 스터디 생성 (블록체인 + DB)
	study, txHash, err := service.CreateStudy(currentUser.ID, req)
	if err != nil {
		fmt.Printf("스터디 생성 실패: %v\n", err)
		http.Error(w, "스터디 생성 실패: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(types.StudyCreateResponse{
		Success:      true,
		StudyID:      study.ID,
		ProxyAddress: study.ProxyAddress,
		TxHash:       txHash,
	})
}
