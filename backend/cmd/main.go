package main

import (
	_ "github.com/lib/pq"
	"github.com/yoshihiro-shu/tech-blog-backend/backend/infrastructure/persistence/cache"
	"github.com/yoshihiro-shu/tech-blog-backend/backend/interfaces/api/server"
	"github.com/yoshihiro-shu/tech-blog-backend/backend/interfaces/model"
	"github.com/yoshihiro-shu/tech-blog-backend/backend/internal/auth"
	"github.com/yoshihiro-shu/tech-blog-backend/backend/internal/config"
	"github.com/yoshihiro-shu/tech-blog-backend/backend/internal/logger"
	"go.uber.org/zap"
)

func main() {
	logger := logger.New()
	conf, err := config.New()
	if err != nil {
		logger.Fatal("failed at load config.", zap.Error(err))
		return
	}
	db, err := model.New(conf)
	if err != nil {
		logger.Fatal("failed at connect Postgres DB.", zap.Error(err))
		return
	}
	cache := cache.New(conf.CacheRedis)

	defer db.Close()

	auth.Init(conf.AccessToken, conf.RefreshToken)
	s := server.New(conf, logger, db, cache)

	s.SetRouters()

	s.Start()
}
