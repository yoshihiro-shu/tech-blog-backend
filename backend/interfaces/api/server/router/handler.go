package router

import (
	"log"
	"net/http"
)

type appHandler func(w http.ResponseWriter, r *http.Request) error

func (h appHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := h(w, r)
	if err != nil {
		switch e := err.(type) {
		case appError:
			// We can retrieve the status here and write out a specific
			// HTTP status code.
			log.Printf("HTTP %d - %s", e.Status, e)
			http.Error(w, e.Error(), e.Status)
		default:
			// Any error types we don't specifically look out for default
			// to serving a HTTP 500
			http.Error(w, http.StatusText(http.StatusInternalServerError),
				http.StatusInternalServerError)
		}
	}
}

type appError struct {
	Status int
	Err    error
}

func (e appError) Error() string {
	return e.Err.Error()
}
