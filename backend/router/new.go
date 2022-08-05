package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yoshi429/test/config"
	"github.com/yoshi429/test/request"
)

type Router struct {
	*mux.Router
	Context *request.Context
}

func New(conf config.Configs) *Router {
	return &Router{
		Router:  mux.NewRouter(),
		Context: request.NewContext(conf),
	}
}

func (r Router) Group(path string) Router {
	r.Router = r.PathPrefix(path).Subrouter()
	return r
}

func (r Router) AppHandle(path string, fn func(http.ResponseWriter, *http.Request) error) *mux.Route {
	return r.Handle(path, AppHandler(fn))
}
