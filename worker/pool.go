package worker

import (
	"baycommit-go/service"
	"baycommit-go/types"
	"fmt"
	"log"
)

// 작업 구조체
type CommitJob struct {
	Commit types.Commit
}

// 전역 채널 (작업 큐)
var JobQueue chan CommitJob

// Worker Pool 시작
func StartWorkerPool(numWorkers int, queueSize int) {
	JobQueue = make(chan CommitJob, queueSize)

	log.Printf("Starting %d workers...\n", numWorkers)

	// Worker 고루틴 N개 생성
	for i := 1; i <= numWorkers; i++ {
		go worker(i)
	}
}

// Worker 함수 (채널에서 작업 꺼내서 처리)
func worker(id int) {
	for job := range JobQueue {
		// [메트릭] Worker Pool 큐 사이즈 감소
		service.WorkerQueueSize.Dec()
		log.Printf("Worker %d processing commit: %s\n", id, job.Commit.Id[:8])

		err := service.ProcessCommit(job.Commit)
		if err != nil {
			fmt.Printf("Worker %d: Commit processing failed: %v\n", id, err)
		} else {
			log.Printf("Worker %d: Commit processed successfully\n", id)
		}
	}
}
