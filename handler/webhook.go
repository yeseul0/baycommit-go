// webhook 핸들러
package handler

import (
	"baycommit-go/service"
	"baycommit-go/types"
	"baycommit-go/worker"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

/*
깃헙 웹훅 verify, json 파싱, commit 있으면
-> service.ProcessCommit() 호출
*/
func WebhookHandler(w http.ResponseWriter, r *http.Request) {
	//Post 요청만 받음
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	body, _ := io.ReadAll(r.Body)
	signature := r.Header.Get("X-Hub-Signature-256")
	githubEvent := r.Header.Get("X-GitHub-Event")
	fmt.Printf("Github webhook received: %s\n", githubEvent)

	//sig 검증
	if !service.VerifySignature(body, signature) {
		fmt.Println("Invalid webhook signature")
		sendResponse(w, false, "Invalid signature")
		return
	}

	//json parsing
	var payload types.GitHubWebhookPayload
	json.Unmarshal(body, &payload)

	//ping 이벤트 처리 (github에 웹훅 등록할때 필요)
	if payload.Zen != "" {
		fmt.Println("Github ping event received")
		sendResponse(w, true, "Ping received successfully")
		return
	}

	//push 이벤트 + 커밋 있을 때만 처리
	if githubEvent == "push" && len(payload.Commits) > 0 {
		sendResponse(w, true, "Webhook received, processing commits...")

		/*
			go func() { // webhook 마다 고루틴 (worker pool로 처리량 제어)
				for _, commit := range payload.Commits {
					commit.RepositoryUrl = payload.Repository.HtmlUrl
					err := service.ProcessCommit(commit)
					if err != nil {
						fmt.Printf("Commit processing failed: %v\n", err)
					}
				}
				fmt.Printf("Successfully processed %d commits in background\n", len(payload.Commits))
			}()
		*/
		for _, commit := range payload.Commits {
			commit.RepositoryUrl = payload.Repository.HtmlUrl
			worker.JobQueue <- worker.CommitJob{Commit: commit} // 작업 구조체로 바꿔서 채널에 넣기만!
		}
		return
	}
	sendResponse(w, true, githubEvent+" event ignored")
}

func sendResponse(w http.ResponseWriter, success bool, message string) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(types.WebhookResponse{
		Success: success,
		Message: message,
	})
}
