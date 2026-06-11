package main

import (
	handler "baycommit-go/handler"
	"fmt"
	"log"
	"net/http"
	"os"

	"baycommit-go/service"
	"baycommit-go/worker"

	"github.com/joho/godotenv"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {

	// .env 파일 로드
	err := godotenv.Load()
	if err != nil {
		log.Println(".env 파일 없음")
	}

	//DB 연결
	err = service.ConnectDB()
	if err != nil {
		log.Fatal("DB 연결 실패:", err)
	}
	fmt.Println("DB 연결됨")

	//블록체인 연결
	err = service.ConnectBlockchain()
	if err != nil {
		log.Fatal("블록체인 연결 실패:", err)
	}
	fmt.Println("블록체인 연결됨")

	//Worker Pool 시작 (워커 10개, 큐 사이즈 100)
	worker.StartWorkerPool(10, 100)

	// 메트릭 초기화
	service.InitMetrics()

	//api url 등록
	http.HandleFunc("/webhook", handler.WebhookHandler)
	http.HandleFunc("/auth/github/callback", handler.GithubCallbackHandler)
	http.HandleFunc("/github/repositories", handler.GithubRepositoriesHandler)
	http.HandleFunc("/study/list", handler.StudyListHandler)
	http.HandleFunc("/study/create", handler.StudyCreateHandler)
	http.HandleFunc("/study/all/commits/today", handler.TodayCommitsHandler)
	http.HandleFunc("/study/repository/register", handler.StudyRepositoryRegisterHandler)
	http.HandleFunc("/study/", handler.StudyRepositoriesHandler)
	http.HandleFunc("/", handler.HealthHandler)
	http.Handle("/metrics", promhttp.Handler()) // Prometheus 메트릭 엔드포인트

	selfURL := os.Getenv("SELF_URL")
	if selfURL != "" {
		handler.StartSelfPing(selfURL)
	}

	//listning
	fmt.Println("서버시작 :8080 listening")
	http.ListenAndServe(":8080", nil)
}
