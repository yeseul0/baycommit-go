package types

type SessionStatus string

// 세션 상태 enum
const (
	StatusPending SessionStatus = "PENDING"
	StatusActive  SessionStatus = "ACTIVE"
	StatusClosed  SessionStatus = "CLOSED"
	StatusFailed  SessionStatus = "FAILED"
)
