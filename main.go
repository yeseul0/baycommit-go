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

	//api url 등록
	http.HandleFunc("/webhook", handler.WebhookHandler) // /webhook 으로 들어오는 모든 메소드 다 받음 ㅋ!
	http.HandleFunc("/", handler.HealthHandler)

	selfURL := os.Getenv("SELF_URL")
	if selfURL != "" {
		handler.StartSelfPing(selfURL)
	}

	//listning
	fmt.Println("서버시작 :8080 listening")
	http.ListenAndServe(":8080", nil)
}
