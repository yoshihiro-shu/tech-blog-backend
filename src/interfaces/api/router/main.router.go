package router

import (
	"github.com/gorilla/mux"
	"github.com/yoshihiro-shu/tech-blog-backend/src/infrastructure/persistence/cache"
	"github.com/yoshihiro-shu/tech-blog-backend/src/interfaces/api/middlewares"
	"github.com/yoshihiro-shu/tech-blog-backend/src/internal/config"
	"github.com/yoshihiro-shu/tech-blog-backend/src/internal/logger"

	"gorm.io/gorm"
)

type MainRouter struct {
	redis     cache.RedisClient
	masterdDB func() *gorm.DB
	repricaDB func() *gorm.DB
	logger    logger.Logger
	conf      config.Configs
}

func NewMainAPI(redis cache.RedisClient, masterdDB func() *gorm.DB, repricaDB func() *gorm.DB, logger logger.Logger, conf config.Configs) *MainRouter {
	return &MainRouter{
		redis:     redis,
		masterdDB: masterdDB,
		repricaDB: repricaDB,
		logger:    logger,
	}
}

func (m *MainRouter) SetRouters(router *mux.Router) {
	router.Use(middlewares.Logger(m.logger))
	router.Use(middlewares.CsrfProtecter(m.conf, m.logger))
	router.Use(middlewares.SetterCsrfToken)
	router.Use(middlewares.Cors(m.conf.Frontend))

	metricsR := &metricsRouter{}
	metricsR.SetRouters(router)

	apiv1 := router.PathPrefix("/api").Subrouter()
	articleR := &articleRouter{
		redis:     m.redis,
		masterdDB: m.masterdDB,
		repricaDB: m.repricaDB,
		logger:    m.logger,
	}
	articleR.SetRouters(apiv1)

	topPageR := &topPageRouter{
		redis:     m.redis,
		masterdDB: m.masterdDB,
		repricaDB: m.repricaDB,
		logger:    m.logger,
	}
	topPageR.SetRouters(apiv1)

	latestArticlesR := &latestArtilesRouter{
		redis:     m.redis,
		masterdDB: m.masterdDB,
		repricaDB: m.repricaDB,
		logger:    m.logger,
	}
	latestArticlesR.SetRouters(apiv1)
}
