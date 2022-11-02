package api

import (
	"net/http"

	"github.com/yoshihiro-shu/draft-backend/interfaces/api/server/request"
	article_cache "github.com/yoshihiro-shu/draft-backend/model/article/cache"
	article_linkages_to_many "github.com/yoshihiro-shu/draft-backend/model/article/linkages/to/many"
)

type articleHandler struct {
	C *request.Context
}

func (a articleHandler) GetTopPage(w http.ResponseWriter, r *http.Request) error {

	var articles []article_linkages_to_many.Article
	err := a.C.Cache.GET(article_cache.TopPageAritcleListKey, &articles)
	if err == nil {
		return a.C.JSON(w, http.StatusOK, articles)
	}

	err = article_linkages_to_many.GetArticleList(a.C.DB(), &articles)
	if err != nil {
		return a.C.JSON(w, http.StatusInternalServerError, err.Error())
	}

	err = a.C.Cache.SET(article_cache.TopPageAritcleListKey, articles)
	return a.C.JSON(w, http.StatusOK, articles)
}

func NewArticleHandler(c *request.Context) *articleHandler {
	return &articleHandler{
		C: c,
	}
}
