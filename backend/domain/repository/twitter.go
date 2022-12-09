package repository

import (
	"github.com/yoshihiro-shu/draft-backend/domain/model"
	"github.com/yoshihiro-shu/draft-backend/internal/config"
)

type TwitterRepository interface {
	GetTimelines(conf config.Configs) (*model.TwitterTimeLine, error)
}
