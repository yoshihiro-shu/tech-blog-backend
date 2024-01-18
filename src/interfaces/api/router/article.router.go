package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yoshihiro-shu/tech-blog-backend/src/infrastructure/persistence/cache"
	"github.com/yoshihiro-shu/tech-blog-backend/src/registory"
	"gorm.io/gorm"
)

type articleRouter struct {
	redis     cache.RedisClient
	masterdDB func() *gorm.DB
	repricaDB func() *gorm.DB
}

func (a *articleRouter) SetRouters(router *mux.Router) {
	h := registory.NewArticleRegistory(
		nil,
		a.masterdDB,
		a.repricaDB,
	)
	article := router.PathPrefix("/articles").Subrouter()
	article.Handle("/{id:[0-9]+}", appHandler(h.Get)).Methods(http.MethodGet)
	article.Handle("/category/{slug}", appHandler(h.GetArticlesByCategory)).Methods(http.MethodGet)
	article.Handle("/tag/{slug}", appHandler(h.GetArticlesByTag)).Methods(http.MethodGet)
}
