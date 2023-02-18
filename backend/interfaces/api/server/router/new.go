package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yoshihiro-shu/draft-backend/interfaces/api/server/httputils"
)

type Router interface {
	Group(path string) Router
	Use(fn ...mux.MiddlewareFunc)
	ServeHTTP(rw http.ResponseWriter, req *http.Request)
	GET(path string, fn func(http.ResponseWriter, *http.Request) error)
	POST(path string, fn func(http.ResponseWriter, *http.Request) error)
	PUT(path string, fn func(http.ResponseWriter, *http.Request) error)
	DELETE(path string, fn func(http.ResponseWriter, *http.Request) error)
}

type router struct {
	Router *mux.Router
}

// type MiddlewareFunc func(http.Handler) http.Handler
type MiddlewareFunc mux.MiddlewareFunc

func New() Router {
	return &router{
		Router: mux.NewRouter(),
	}
}

func (r router) Group(path string) Router {
	r.Router = r.Router.PathPrefix(path).Subrouter()
	return r
}

func (r router) Use(fn ...mux.MiddlewareFunc) {
	r.Router.Use(fn...)
}

func (r router) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	r.Router.ServeHTTP(rw, req)
}

func (r router) GET(path string, fn func(http.ResponseWriter, *http.Request) error) {
	r.Router.Handle(path, httputils.Handler(fn)).Methods(http.MethodGet)
}

func (r router) POST(path string, fn func(http.ResponseWriter, *http.Request) error) {
	r.Router.Handle(path, httputils.Handler(fn)).Methods(http.MethodPost)
}

func (r router) PUT(path string, fn func(http.ResponseWriter, *http.Request) error) {
	r.Router.Handle(path, httputils.Handler(fn)).Methods(http.MethodPut)
}

func (r router) DELETE(path string, fn func(http.ResponseWriter, *http.Request) error) {
	r.Router.Handle(path, httputils.Handler(fn)).Methods(http.MethodDelete)
}
