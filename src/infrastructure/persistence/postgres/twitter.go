package postgres

import (
	"encoding/json"

	"github.com/yoshihiro-shu/tech-blog-backend/src/domain/model"
	"github.com/yoshihiro-shu/tech-blog-backend/src/domain/repository"
	"github.com/yoshihiro-shu/tech-blog-backend/src/internal/config"
	"github.com/yoshihiro-shu/tech-blog-backend/src/internal/twitter_api"
)

type twitterPersistence struct {
}

func NewTwitterPersistence() repository.TwitterRepository {
	return &twitterPersistence{}
}

func (tp *twitterPersistence) GetTimelines(conf config.Configs) (*model.TwitterTimeLine, error) {
	var resTimeline model.TwitterTimeLine
	twitter := twitter_api.NewClient(conf)
	b, err := twitter.GetTimeLines()
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &resTimeline)
	if err != nil {
		return nil, err
	}

	return &resTimeline, nil
}
