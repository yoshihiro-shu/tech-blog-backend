package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yoshihiro-shu/tech-blog-backend/src/infrastructure/persistence/cache"
	"github.com/yoshihiro-shu/tech-blog-backend/src/internal/logger"
	"github.com/yoshihiro-shu/tech-blog-backend/src/registory"
	"gorm.io/gorm"
)

type latestArtilesRouter struct {
	redis     cache.RedisClient
	masterdDB func() *gorm.DB
	repricaDB func() *gorm.DB
	logger    logger.Logger
}

func (r *latestArtilesRouter) SetRouters(router *mux.Router) {
	latestArticlesHandler := registory.NewLatestArticlesRegistory(
		r.redis,
		r.logger,
		r.masterdDB,
		r.repricaDB,
	)
	latest := router.PathPrefix("/new").Subrouter()

	latest.Handle("/{page:[0-9]+}", nil).Methods(http.MethodOptions)
	latest.Handle("/{page:[0-9]+}", appHandler(latestArticlesHandler.Get)).Methods(http.MethodGet)
}
