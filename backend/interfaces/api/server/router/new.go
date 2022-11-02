package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yoshihiro-shu/draft-backend/interfaces/api/server/httputils"
	"github.com/yoshihiro-shu/draft-backend/internal/config"
)

type Router struct {
	*mux.Router
	Config config.Configs
}

func New(conf config.Configs) *Router {
	return &Router{
		Router: mux.NewRouter(),
		Config: conf,
	}
}

func (r Router) Group(path string) Router {
	r.Router = r.PathPrefix(path).Subrouter()
	return r
}

func (r Router) AppHandle(path string, fn func(http.ResponseWriter, *http.Request) error) *mux.Route {
	return r.Handle(path, httputils.AppHandlerFunc(fn))
}
