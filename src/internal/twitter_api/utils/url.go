package twitter_utils

import (
	"fmt"
	"net/url"
)

const twitterApiV2 = "https://api.twitter.com"

func CreateUrlApiV2(urlFormat string, args string) (*url.URL, error) {
	endpoint := fmt.Sprintf(urlFormat, args)

	endpoint = twitterApiV2 + endpoint

	url, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}

	return url, nil
}

func SetParams(url *url.URL, params map[string]string) {
	q := url.Query()
	for i, v := range params {
		q.Set(i, v)
	}

	url.RawQuery = q.Encode()
}
