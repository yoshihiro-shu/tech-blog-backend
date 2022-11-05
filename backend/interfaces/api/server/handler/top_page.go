package handler

import (
	"net/http"

	"github.com/yoshihiro-shu/draft-backend/application/usecase"
	"github.com/yoshihiro-shu/draft-backend/domain/model"
	"github.com/yoshihiro-shu/draft-backend/interfaces/api/server/request"
	article_cache "github.com/yoshihiro-shu/draft-backend/internal/model/article/cache"
)

type TopPageHandler interface {
	Get(w http.ResponseWriter, r *http.Request) error
}

type topPageHandler struct {
	topPageUseCase usecase.TopPageUseCase
	C              *request.Context
}

func NewTopPageHandler(topPageUseCase usecase.TopPageUseCase, c *request.Context) *topPageHandler {
	return &topPageHandler{
		topPageUseCase: topPageUseCase,
		C:              c,
	}
}

type responseTopPage struct {
	Article []model.Article `json:"Article"`
}

func (tp topPageHandler) Get(w http.ResponseWriter, r *http.Request) error {
	var articles []model.Article

	err := tp.C.Cache.GET(article_cache.TopPageAritcleListKey, &articles)
	if err == nil {
		return tp.C.JSON(w, http.StatusOK, articles)
	}

	err = tp.topPageUseCase.GetArticles(&articles)
	if err != nil {
		return tp.C.JSON(w, http.StatusInternalServerError, err.Error())
	}

	// res := &responseTopPage{
	// 	Article: articles,
	// }

	_ = tp.C.Cache.SET(article_cache.TopPageAritcleListKey, articles)
	return tp.C.JSON(w, http.StatusOK, articles)
}
