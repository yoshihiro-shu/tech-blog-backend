package tweet_timeline

import (
	twitter_utils "github.com/yoshihiro-shu/draft-backend/internal/twitter_api/utils"
)

const apiEndpoint = "/2/users/%s/tweets"

var tweetParams = map[string]string{"tweet.fields": "created_at"}

func Do(userId, bearerToken string) ([]byte, error) {
	url, _ := twitter_utils.CreateUrlApiV2(apiEndpoint, userId)

	twitter_utils.SetParams(url, tweetParams)

	data, err := twitter_utils.ConnectToEndpointHttpGet(url.String(), bearerToken)
	if err != nil {
		return nil, err
	}
	return data, nil
}
