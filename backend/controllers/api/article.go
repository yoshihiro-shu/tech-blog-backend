package api

import (
	"net/http"

	article_linkages_to_many "github.com/yoshihiro-shu/draft-backend/model/article/linkages/to/many"
	"github.com/yoshihiro-shu/draft-backend/request"
)

type articleHandler struct {
	C *request.Context
}

func (a articleHandler) GetTopPage(w http.ResponseWriter, r *http.Request) error {

	ac := &article_linkages_to_many.Article{}

	articles, err := ac.GetArticleList(a.C.DB())
	if err != nil {
		return a.C.JSON(w, http.StatusInternalServerError, err.Error())
	}
	return a.C.JSON(w, http.StatusOK, articles)
}

func NewArticleHandler(c *request.Context) *articleHandler {
	return &articleHandler{
		C: c,
	}
}
