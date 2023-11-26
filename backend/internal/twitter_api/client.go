package twitter_api

import (
	"fmt"

	"github.com/yoshihiro-shu/tech-blog-backend/backend/internal/config"
	tweet_timeline "github.com/yoshihiro-shu/tech-blog-backend/backend/internal/twitter_api/tweet/timeline"
)

type TwitterClient struct {
	userId       string
	apiKey       string
	apiKeySecret string
	bearerToken  string
}

func NewClient(conf config.Configs) *TwitterClient {
	return &TwitterClient{
		userId:       conf.Twitter.UserId,
		apiKey:       conf.Twitter.Apikey,
		apiKeySecret: conf.Twitter.ApiKeySecret,
		bearerToken:  conf.Twitter.BearerToken,
	}
}

func (t TwitterClient) GetBearerToken() string {
	return fmt.Sprintf("Bearer %s", t.bearerToken)
}

func (t TwitterClient) GetTimeLines() ([]byte, error) {
	return tweet_timeline.Do(t.userId, t.GetBearerToken())
}
