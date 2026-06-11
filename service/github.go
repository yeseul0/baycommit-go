package service

import (
	"baycommit-go/types"
	"encoding/json"
	"fmt"
	"net/http"
)

// GitHub API 레포 응답 (내부용 - snake_case)
type githubRepoRaw struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	FullName    string `json:"full_name"`
	HtmlUrl     string `json:"html_url"`
	Description string `json:"description"`
	Private     bool   `json:"private"`
	UpdatedAt   string `json:"updated_at"`
	Language    string `json:"language"`
	Fork        bool   `json:"fork"`
}

// GitHub access token으로 유저 레포 목록 조회
func GetUserRepositories(githubAccessToken string) ([]types.GitHubRepository, error) {
	req, _ := http.NewRequest("GET", "https://api.github.com/user/repos?per_page=100&sort=updated", nil)
	req.Header.Set("Authorization", "Bearer "+githubAccessToken)
	req.Header.Set("Accept", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("GitHub 레포 조회 실패: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusUnauthorized {
		return nil, fmt.Errorf("GitHub 토큰이 만료되었습니다. 다시 로그인하세요.")
	}

	var raw []githubRepoRaw
	if err := json.NewDecoder(resp.Body).Decode(&raw); err != nil {
		return nil, fmt.Errorf("응답 파싱 실패: %v", err)
	}

	repos := make([]types.GitHubRepository, 0, len(raw))
	for _, r := range raw {
		repos = append(repos, types.GitHubRepository{
			ID:          r.ID,
			Name:        r.Name,
			FullName:    r.FullName,
			HtmlUrl:     r.HtmlUrl,
			Description: r.Description,
			Private:     r.Private,
			UpdatedAt:   r.UpdatedAt,
			Language:    r.Language,
			Fork:        r.Fork,
		})
	}

	return repos, nil
}
