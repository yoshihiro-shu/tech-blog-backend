package request

import (
	"context"

	"github.com/go-pg/pg"
	"github.com/go-playground/validator/v10"
	"github.com/yoshihiro-shu/draft-backend/backend/interfaces/api/server/auth"
	"github.com/yoshihiro-shu/draft-backend/backend/interfaces/api/server/cache"
	"github.com/yoshihiro-shu/draft-backend/backend/interfaces/api/server/model"
	"github.com/yoshihiro-shu/draft-backend/backend/internal/config"
	"github.com/yoshihiro-shu/draft-backend/backend/internal/pkg/logger"
	"gorm.io/gorm"
)

type Context struct {
	db       *model.DBContext
	cache    cache.RedisClient
	Conf     config.Configs
	Logger   logger.Logger
	validate *validator.Validate
}

func NewContext(conf config.Configs, logger logger.Logger, db *model.DBContext, cache cache.RedisClient) *Context {
	return &Context{
		db:       db,
		cache:    cache,
		Conf:     conf,
		Logger:   logger,
		validate: validator.New(),
	}
}

func (c Context) MasterDB() *pg.DB {
	return c.db.Master()
}

func (c Context) RepricaDB() *pg.DB {
	return c.db.Reprica()
}

func (c Context) DBPrimary() *gorm.DB {
	return c.db.Primary()
}

func (c Context) Cache() cache.RedisClient {
	return c.cache
}

func (c Context) GetAuthUserID(ctx context.Context) interface{} {
	return ctx.Value(auth.UserKey)
}
