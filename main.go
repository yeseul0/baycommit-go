package main

import (
	handler "baycommit-go/Handler"
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
	// .env 파일 로드
	err := godotenv.Load()
	if err != nil {
		log.Println(".env 파일 없음")
	}

	http.HandleFunc("/webhook", handler.WebhookHandler) // /webhook 으로 들어오는 모든 메소드 다 받음 ㅋ!

	fmt.Println("서버시작: http://localhost:8080:")
	http.ListenAndServe(":8080", nil)
}
