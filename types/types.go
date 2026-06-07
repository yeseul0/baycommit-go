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
