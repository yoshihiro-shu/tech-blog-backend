package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yoshihiro-shu/tech-blog-backend/src/infrastructure/persistence/cache"
	"github.com/yoshihiro-shu/tech-blog-backend/src/internal/logger"
	"github.com/yoshihiro-shu/tech-blog-backend/src/registory"
)

type profileRouter struct {
	token  string
	redis  cache.RedisClient
	logger logger.Logger
}

func (p *profileRouter) SetRouters(router *mux.Router) {
	h := registory.NewProfileRegistory(
		p.token,
		p.redis,
		p.logger,
	)

	profile := router.PathPrefix("/profile").Subrouter()
	profile.Handle("/resume", nil).Methods(http.MethodOptions)
	profile.Handle("/resume", appHandler(h.GetResume)).Methods(http.MethodGet)
}
