package request

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/yoshihiro-shu/tech-blog-backend/src/infrastructure/persistence/cache"
	"github.com/yoshihiro-shu/tech-blog-backend/src/interfaces/model"
	"github.com/yoshihiro-shu/tech-blog-backend/src/internal/auth"
	"github.com/yoshihiro-shu/tech-blog-backend/src/internal/config"
	"github.com/yoshihiro-shu/tech-blog-backend/src/internal/logger"
	"gorm.io/gorm"
)

type Context struct {
	db       model.DBClient
	cache    cache.RedisClient
	Conf     config.Configs
	Logger   logger.Logger
	validate *validator.Validate
}

func NewContext(conf config.Configs, logger logger.Logger, db model.DBClient, cache cache.RedisClient) *Context {
	return &Context{
		db:       db,
		cache:    cache,
		Conf:     conf,
		Logger:   logger,
		validate: validator.New(),
	}
}

func (c Context) MasterDB() *gorm.DB {
	return c.db.Master()
}

func (c Context) RepricaDB() *gorm.DB {
	return c.db.Reprica()
}

func (c Context) Cache() cache.RedisClient {
	return c.cache
}

func (c Context) GetAuthUserID(ctx context.Context) interface{} {
	return ctx.Value(auth.UserKey)
}
