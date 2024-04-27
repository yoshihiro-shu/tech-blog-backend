package github_api

import (
	"fmt"
	"io"
	"net/http"
)

type client struct {
	owner string
	repo  string
	token string
}

func NewClient(owner, repo, token string) *client {
	return &client{
		owner: owner,
		repo:  repo,
		token: token,
	}
}

// GetRepositoryContent is get github repository content by path
func (c *client) GetRepositoryContent(path string) ([]byte, error) {
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/contents/%s", c.owner, c.repo, path)

	// http.Clientの生成
	client := &http.Client{}

	// http.NewRequestでGETリクエストを作成
	// curl -H "Authorization: Bearer ${Github Personal Token}" -H "Accept: application/vnd.github.html+json" https://api.github.com/repos/yoshihiro-shu/Resume/contents/README.md
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// 個人アクセストークンをAuthorizationヘッダーに設定
	// YOUR_TOKENには実際のアクセストークンを設定してください
	bearerToken := fmt.Sprintf("Bearer %s", c.token)
	req.Header.Add("Authorization", bearerToken)

	// https://docs.github.com/ja/rest/repos/contents?apiVersion=2022-11-28#get-repository-content
	req.Header.Add("Accept", "application/vnd.github.raw+json")

	// リクエストの送信
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}
