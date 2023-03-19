package main

import (
	_ "github.com/lib/pq"
	"github.com/yoshihiro-shu/draft-backend/backend/interfaces/api/server"
	"github.com/yoshihiro-shu/draft-backend/backend/interfaces/api/server/auth"
	"github.com/yoshihiro-shu/draft-backend/backend/interfaces/api/server/cache"
	"github.com/yoshihiro-shu/draft-backend/backend/interfaces/api/server/model"
	"github.com/yoshihiro-shu/draft-backend/backend/internal/config"
	"github.com/yoshihiro-shu/draft-backend/backend/internal/pkg/logger"
	"go.uber.org/zap"
)

func main() {
	logger := logger.New()
	conf, err := config.New()
	if err != nil {
		logger.Fatal("failed at load config.", zap.Error(err))
		return
	}
	db := model.New(conf)
	cache := cache.New(conf.CacheRedis)

	defer db.Close()

	auth.Init(conf.AccessToken, conf.RefreshToken)
	s := server.New(conf, logger, db, cache)

	s.SetRouters()

	s.Start()
}
