package repository

import (
	"github.com/yoshihiro-shu/tech-blog-backend/src/domain/model"
	"github.com/yoshihiro-shu/tech-blog-backend/src/internal/config"
)

type TwitterRepository interface {
	GetTimelines(conf config.Configs) (*model.TwitterTimeLine, error)
}
