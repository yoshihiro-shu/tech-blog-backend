package model

import "time"

type (
	TwitterTimeLine struct {
		Data []Tweet     `json:"data"`
		Meta TwitterMeta `json:"meta"`
	}
	Tweet struct {
		EditHistoryTweetIds []string  `json:"edit_history_tweet_ids"`
		Text                string    `json:"text"`
		CreatedAt           time.Time `json:"created_at"`
		ID                  string    `json:"id"`
	}
	TwitterMeta struct {
		NextToken   string `json:"next_token"`
		ResultCount int    `json:"result_count"`
		NewestID    string `json:"newest_id"`
		OldestID    string `json:"oldest_id"`
	}
)
