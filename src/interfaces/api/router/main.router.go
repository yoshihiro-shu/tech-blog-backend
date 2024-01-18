package router

import (
	"github.com/gorilla/mux"
	"github.com/yoshihiro-shu/tech-blog-backend/src/interfaces/api/handler"
)

type MainRouter struct {
	articleR        *articleRouter
	topPageR        *topPageRouter
	latestArticlesR *latestArtilesRouter
}

func NewMainAPI(a *articleRouter, t *topPageRouter, l *latestArtilesRouter) *MainRouter {
	return &MainRouter{articleR: a, topPageR: t, latestArticlesR: l}
}

func (m *MainRouter) SetRouters(h handler.UserHandler) {
	router := mux.NewRouter()
	m.articleR.SetRouters(router)
	m.topPageR.SetRouters(router)
	m.latestArticlesR.SetRouters(router)
}
