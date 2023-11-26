package router

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Router interface {
	Group(path string) Router
	Use(fn ...func(http.Handler) http.Handler)
	ServeHTTP(rw http.ResponseWriter, req *http.Request)
	Handle(path string, handler http.Handler)
	GET(path string, fn func(http.ResponseWriter, *http.Request) error)
	POST(path string, fn func(http.ResponseWriter, *http.Request) error)
	PUT(path string, fn func(http.ResponseWriter, *http.Request) error)
	DELETE(path string, fn func(http.ResponseWriter, *http.Request) error)
}

type router struct {
	Router *mux.Router
}

type MiddlewareFunc func(http.Handler) http.Handler

func New() Router {
	return &router{
		Router: mux.NewRouter(),
	}
}

func (r router) Handle(path string, handler http.Handler) {
	r.Router.Handle(path, handler)
}

func (r router) Group(path string) Router {
	r.Router = r.Router.PathPrefix(path).Subrouter()
	return r
}

func (r router) Use(fns ...func(http.Handler) http.Handler) {

	middlewareFuncs := make([]mux.MiddlewareFunc, len(fns))
	for i, v := range fns {
		middlewareFuncs[i] = mux.MiddlewareFunc(v)
	}

	r.Router.Use(middlewareFuncs...)
}

func (r router) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	r.Router.ServeHTTP(rw, req)
}

func (r router) GET(path string, fn func(http.ResponseWriter, *http.Request) error) {
	r.Router.Handle(path, appHandler(fn)).Methods(http.MethodGet, http.MethodOptions)
}

func (r router) POST(path string, fn func(http.ResponseWriter, *http.Request) error) {
	r.Router.Handle(path, appHandler(fn)).Methods(http.MethodPost, http.MethodOptions)
}

func (r router) PUT(path string, fn func(http.ResponseWriter, *http.Request) error) {
	r.Router.Handle(path, appHandler(fn)).Methods(http.MethodPut, http.MethodOptions)
}

func (r router) DELETE(path string, fn func(http.ResponseWriter, *http.Request) error) {
	r.Router.Handle(path, appHandler(fn)).Methods(http.MethodDelete, http.MethodOptions)
}
