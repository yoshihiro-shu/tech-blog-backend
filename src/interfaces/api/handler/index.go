package handler

import (
	"net/http"

	"github.com/yoshihiro-shu/tech-blog-backend/src/interfaces/api/request"
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
func NewIndexHandler(c *request.Context) *indexHandler {
	return &indexHandler{c}
}
