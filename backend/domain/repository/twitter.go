package repository

import (
	"github.com/yoshihiro-shu/draft-backend/backend/domain/model"
	"github.com/yoshihiro-shu/draft-backend/backend/internal/config"
)

type TwitterRepository interface {
	GetTimelines(conf config.Configs) (*model.TwitterTimeLine, error)
}
