package twitter_utils

import (
	"io"
	"net/http"
	"time"
)

// request client for http method is GET
func ConnectToEndpointHttpGet(url, bearerToken string) ([]byte, error) {
	clinet := &http.Client{
		Timeout: time.Second * 10,
	}
	req, _ := http.NewRequest(http.MethodGet, url, nil)

	setBearerOauth(req, bearerToken)

	res, err := clinet.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	b, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, nil
	}

	return b, nil
}
