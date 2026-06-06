package service

import "github.com/prometheus/client_golang/prometheus"

var (
	// 웹훅 수신 횟수
	WebhookTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "baycommit_webhook_total",
			Help: "웹훅 수신 횟수",
		},
		[]string{"status"}, // status: success | error
	)

	// 커밋 처리 횟수
	CommitProcessTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "baycommit_commit_process_total",
			Help: "커밋 처리 횟수",
		},
		[]string{"result"}, // result: success | skip | error
	)

	// DB 쿼리 소요시간
	DBQueryDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "baycommit_db_query_duration_seconds",
			Help:    "DB 쿼리 소요시간",
			Buckets: []float64{0.001, 0.005, 0.01, 0.05, 0.1, 0.5},
		},
		[]string{"query"}, // query: get_studies_by_repo | get_wallet_address | create_session | update_session
	)

	// 블록체인 RPC 소요시간
	BlockchainRPCDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "baycommit_blockchain_rpc_duration_seconds",
			Help:    "블록체인 RPC 소요시간",
			Buckets: []float64{0.1, 0.5, 1, 2, 5, 10},
		},
		[]string{"method"}, // method: get_study_day_info | get_commit_time | start_today_study | track_commit
	)

	// Worker Pool 큐 사이즈
	WorkerQueueSize = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "baycommit_worker_queue_size",
			Help: "Worker Pool 현재 큐 사이즈",
		},
	)
)

func InitMetrics() {
	prometheus.MustRegister(
		WebhookTotal,
		CommitProcessTotal,
		DBQueryDuration,
		BlockchainRPCDuration,
		WorkerQueueSize,
	)
}
