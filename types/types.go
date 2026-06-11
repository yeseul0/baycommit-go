package types

type GitHubWebhookPayload struct {
	Zen        string   `json:"zen"` // ping 이벤트 확인용
	Commits    []Commit `json:"commits"`
	Repository struct {
		HtmlUrl string `json:"html_url"`
	} `json:"repository"`
}

type Commit struct {
	Id        string `json:"id"`
	Message   string `json:"message"`
	Timestamp string `json:"timestamp"`
	Author    struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	} `json:"author"`
	RepositoryUrl string // 나중에 채워넣기
}

// API 응답
type WebhookResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

// GitHub API /user 응답
type GitHubUser struct {
	ID    int    `json:"id"`
	Login string `json:"login"`
	Email string `json:"email"`
}

// GitHub OAuth access token 응답
type GitHubTokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	Scope       string `json:"scope"`
}

// GitHub /user/emails 응답
type GitHubEmail struct {
	Email    string `json:"email"`
	Primary  bool   `json:"primary"`
	Verified bool   `json:"verified"`
}

// ===== /study/list 응답 타입 =====

type CommitInfo struct {
	CommitId      string `json:"commitId"`
	CommitMessage string `json:"commitMessage"`
	CommitTime    int64  `json:"commitTime"`
}

type ParticipantInfo struct {
	GithubEmail   string      `json:"githubEmail"`
	GithubHandle  string      `json:"githubHandle"`
	WalletAddress string      `json:"walletAddress"`
	Commit        *CommitInfo `json:"commit"` // null이면 커밋 안 함
}

type TodaySession struct {
	SessionId    uint              `json:"sessionId"`
	Status       string            `json:"status"`
	StudyDate    string            `json:"studyDate"`
	StartedAt    int64             `json:"startedAt"`
	Participants []ParticipantInfo `json:"participants"`
}

type StudyListItem struct {
	ID                      uint          `json:"id"`
	StudyName               string        `json:"studyName"`
	CreatedAt               string        `json:"createdAt"`
	ProxyAddress            string        `json:"proxyAddress"`
	StudyStartTime          int64         `json:"studyStartTime"`
	StudyEndTime            int64         `json:"studyEndTime"`
	IsOwner                 bool          `json:"isOwner"`
	IsParticipating         bool          `json:"isParticipating"`
	HasRegisteredRepository bool          `json:"hasRegisteredRepository"`
	MyRepoUrl               string        `json:"myRepoUrl"`
	ParticipantCount        int           `json:"participantCount"`
	TodaySession            *TodaySession `json:"todaySession"` // null이면 오늘 세션 없음
}

type StudyListResponse struct {
	Success bool            `json:"success"`
	Studies []StudyListItem `json:"studies"`
}

// ===== /github/repositories 응답 타입 =====

type GitHubRepository struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	FullName    string `json:"fullName"`
	HtmlUrl     string `json:"htmlUrl"`
	Description string `json:"description"`
	Private     bool   `json:"private"`
	UpdatedAt   string `json:"updatedAt"`
	Language    string `json:"language"`
	Fork        bool   `json:"fork"`
}

type GitHubRepositoriesResponse struct {
	Success      bool                `json:"success"`
	Repositories []GitHubRepository  `json:"repositories"`
}

// ===== /study/all/commits/today 응답 타입 =====

type TodayCommitSession struct {
	SessionId    uint              `json:"sessionId"`
	StudyDate    string            `json:"studyDate"`
	Status       string            `json:"status"`
	StartedAt    int64             `json:"startedAt"`
	Participants []ParticipantInfo `json:"participants"`
}

type TodayCommitStudy struct {
	StudyName    string               `json:"studyName"`
	ProxyAddress string               `json:"proxyAddress"`
	Sessions     []TodayCommitSession `json:"sessions"`
}

type TodayCommitsResponse struct {
	Success bool               `json:"success"`
	Studies []TodayCommitStudy `json:"studies"`
}

// ===== /study/create 요청/응답 타입 =====

type StudyCreateRequest struct {
	StudyName      string `json:"studyName"`
	DepositAmount  string `json:"depositAmount"`  // USDC 단위 (wei), 문자열로 받아 big.Int 변환
	PenaltyAmount  string `json:"penaltyAmount"`  // USDC 단위 (wei), 문자열로 받아 big.Int 변환
	StudyStartTime int64  `json:"studyStartTime"` // Unix timestamp
	StudyEndTime   int64  `json:"studyEndTime"`   // Unix timestamp
}

type StudyCreateResponse struct {
	Success      bool   `json:"success"`
	StudyID      uint   `json:"studyId"`
	ProxyAddress string `json:"proxyAddress"`
	TxHash       string `json:"txHash"`
}
