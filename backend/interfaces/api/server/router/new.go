package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yoshihiro-shu/draft-backend/interfaces/api/server/httputils"
)

type Router struct {
	*mux.Router
}

func New() *Router {
	return &Router{
		Router: mux.NewRouter(),
	}
}

func (r Router) Group(path string) Router {
	r.Router = r.PathPrefix(path).Subrouter()
	return r
}

func (r Router) AppHandle(path string, fn func(http.ResponseWriter, *http.Request) error) *mux.Route {
	return r.Handle(path, httputils.Handler(fn))
}
