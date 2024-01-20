package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yoshihiro-shu/tech-blog-backend/src/infrastructure/persistence/cache"
	"github.com/yoshihiro-shu/tech-blog-backend/src/internal/logger"
	"github.com/yoshihiro-shu/tech-blog-backend/src/registory"
	"gorm.io/gorm"
)

type topPageRouter struct {
	redis     cache.RedisClient
	masterdDB func() *gorm.DB
	repricaDB func() *gorm.DB
	logger    logger.Logger
}

func (r *topPageRouter) SetRouters(router *mux.Router) {
	topPageHandler := registory.NewTopPageRegistory(
		r.redis,
		r.logger,
		r.masterdDB,
		r.repricaDB,
	)

	router.Handle("/top", nil).Methods(http.MethodOptions)
	router.Handle("/top", appHandler(topPageHandler.Get)).Methods(http.MethodGet)
}
