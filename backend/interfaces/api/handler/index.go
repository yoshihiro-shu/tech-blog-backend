package handler

import (
	"net/http"

	"github.com/yoshihiro-shu/draft-backend/backend/interfaces/api/request"
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
func (h indexHandler) AuthIndex(w http.ResponseWriter, r *http.Request) error {
	id := h.GetAuthUserID(r.Context())
	return h.JSON(w, http.StatusOK, id)
}

func NewIndexHandler(c *request.Context) *indexHandler {
	return &indexHandler{c}
}
