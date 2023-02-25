package main

import (
	_ "github.com/lib/pq"
	"github.com/yoshihiro-shu/draft-backend/backend/interfaces/api/server"
	"github.com/yoshihiro-shu/draft-backend/backend/interfaces/api/server/auth"
	"github.com/yoshihiro-shu/draft-backend/backend/interfaces/api/server/cache"
	"github.com/yoshihiro-shu/draft-backend/backend/interfaces/api/server/model"
	"github.com/yoshihiro-shu/draft-backend/backend/internal/config"
	"github.com/yoshihiro-shu/draft-backend/backend/internal/pkg/logger"
)

func main() {
	conf := config.New()
	logger := logger.New()
	db := model.New(conf)
	cache := cache.New(conf.CacheRedis)

	defer db.Close()

	auth.Init(conf.AccessToken, conf.RefreshToken)
	s := server.New(conf, logger, db, cache)

	s.SetRouters()

	s.Start()
}
