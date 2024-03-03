package github_api_test

import (
	"bytes"
	"html/template"
	"os"
	"testing"

	"github.com/yoshihiro-shu/tech-blog-backend/src/internal/github_api"
)

var token = os.Getenv("GITHUB_TOKEN")

func TestGetRepositoryContent(t *testing.T) {
	client := github_api.NewClient("yoshihiro-shu", "Resume", token)

	res, err := client.GetRepositoryContent("README.md")
	if err != nil {
		t.Error(err)
	}

	if !isHTMLTemplate(string(res)) {
		t.Error("Not HTML Template")
	}
}

// isHTMLTemplate checks if a string can be parsed as an HTML template
func isHTMLTemplate(str string) bool {
	tmpl, err := template.New("test").Parse(str)
	if err != nil {
		// 解析中にエラーが発生した場合、おそらくHTMLではない
		return false
	}

	// テンプレートを実行してみてエラーが発生するかチェック
	var buf bytes.Buffer
	err = tmpl.Execute(&buf, nil)
	return err == nil
}
