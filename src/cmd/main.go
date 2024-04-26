package main

import (
	_ "github.com/lib/pq"
	"github.com/yoshihiro-shu/tech-blog-backend/src/infrastructure/persistence/cache"
	"github.com/yoshihiro-shu/tech-blog-backend/src/interfaces/api/server"
	"github.com/yoshihiro-shu/tech-blog-backend/src/interfaces/model"
	"github.com/yoshihiro-shu/tech-blog-backend/src/internal/auth"
	"github.com/yoshihiro-shu/tech-blog-backend/src/internal/config"
	"github.com/yoshihiro-shu/tech-blog-backend/src/internal/logger"
	"go.uber.org/zap"
)

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	logger := logger.New()
	conf, err := config.New()
	logger.Info("config", zap.Any("config", conf))
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
	s := server.New(conf, logger, db.Master, db.Reprica, cache)
	s.Start()
}
