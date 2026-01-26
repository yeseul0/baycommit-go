package handler

import (
	"fmt"
	"net/http"
	"time"
)

func StartSelfPing(url string) {
	go func() {
		for {
			time.Sleep(8 * time.Minute)
			resp, err := http.Get(url)
			if err != nil {
				fmt.Println("셀프 핑 실패:", err)
				continue
			}
			resp.Body.Close()
			fmt.Println("셀프핑!")
		}
	}()
}

// health check
func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
