package handler

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yoshihiro-shu/draft-backend/backend/interfaces/api/server/request"
)

type indexHandler struct {
	*request.Context
}

type TestRedis struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func (h indexHandler) Index(w http.ResponseWriter, r *http.Request) error {
	return h.JSON(w, http.StatusOK, "HELLO WORLD")
}

func (h indexHandler) TestHandler(w http.ResponseWriter, r *http.Request) error {
	fmt.Println("UNKOOOOO")
	fmt.Printf("RequestContext: %#v\n", h.Context)
	fmt.Fprintf(w, "RequestContext: %#v\n", h.Context)
	fmt.Println("UNKOOOOO")
	return h.Context.JSON(w, http.StatusOK, h.Context)
}

func (h indexHandler) AuthIndex(w http.ResponseWriter, r *http.Request) error {
	id := h.GetAuthUserID(r.Context())
	return h.JSON(w, http.StatusOK, id)
}

func (h indexHandler) TestSetRedis(w http.ResponseWriter, r *http.Request) error {
	k := r.FormValue("key")
	v := r.FormValue("value")

	t := TestRedis{
		Key:   k,
		Value: v,
	}

	err := h.Cache().SET(k, t)
	if err != nil {
		return h.JSON(w, http.StatusInternalServerError, err.Error())
	}

	cookie := &http.Cookie{
		Name:  k,
		Value: v,
	}

	http.SetCookie(w, cookie)

	return h.JSON(w, http.StatusOK, t)
}

func (h indexHandler) TestGetRedis(w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)
	k := vars["key"]

	t := TestRedis{}
	err := h.Cache().GET(k, &t)
	if err != nil {
		return h.JSON(w, http.StatusInternalServerError, err.Error())
	}

	return h.JSON(w, http.StatusOK, t)
}

func NewIndexHandler(c *request.Context) *indexHandler {
	return &indexHandler{c}
}
