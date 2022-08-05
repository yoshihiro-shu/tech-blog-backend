package handler

import (
	"fmt"
	"net/http"

	"github.com/yoshi429/test/request"
)

type Handler struct {
	Context *request.Context
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
