package handler

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yoshi429/draft-backend/request"
)

type Handler struct {
	Context *request.Context
}

type TestRedis struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func (h Handler) Index(w http.ResponseWriter, r *http.Request) error {
	return h.Context.JSON(w, http.StatusOK, "HELLO WORLD")
}

func (h Handler) TestHandler(w http.ResponseWriter, r *http.Request) error {
	fmt.Println("UNKOOOOO")
	fmt.Printf("RequestContext: %#v\n", h.Context)
	fmt.Fprintf(w, "RequestContext: %#v\n", h.Context)
	fmt.Println("UNKOOOOO")
	return h.Context.JSON(w, http.StatusOK, h.Context)
}

func (h Handler) AuthIndex(w http.ResponseWriter, r *http.Request) error {
	id := h.Context.GetAuthUserID(r.Context())
	return h.Context.JSON(w, http.StatusOK, id)
}

func (h Handler) TestSetRedis(w http.ResponseWriter, r *http.Request) error {
	k := r.FormValue("key")
	v := r.FormValue("value")

	t := TestRedis{
		Key:   k,
		Value: v,
	}

	err := h.Context.Cache.SET(k, t)
	if err != nil {
		return h.Context.JSON(w, http.StatusInternalServerError, err.Error())
	}

	return h.Context.JSON(w, http.StatusOK, t)
}

func (h Handler) TestGetRedis(w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)
	k := vars["key"]

	t := TestRedis{}
	err := h.Context.Cache.GET(k, &t)
	if err != nil {
		return h.Context.JSON(w, http.StatusInternalServerError, err.Error())
	}

	return h.Context.JSON(w, http.StatusOK, t)
}
