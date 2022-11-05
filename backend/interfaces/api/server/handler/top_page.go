package handler

import (
	"net/http"

	"github.com/yoshihiro-shu/draft-backend/application/usecase"
	"github.com/yoshihiro-shu/draft-backend/domain/model"
	"github.com/yoshihiro-shu/draft-backend/interfaces/api/server/request"
	article_cache "github.com/yoshihiro-shu/draft-backend/internal/model/article/cache"
	"github.com/yoshihiro-shu/draft-backend/internal/pkg/pager"
)

type TopPageHandler interface {
	Get(w http.ResponseWriter, r *http.Request) error
}

type topPageHandler struct {
	topPageUseCase usecase.TopPageUseCase
	C              *request.Context
}

const (
	// 一ページあたりの記事数
	numberOfArticlePer1Page = 1
)

func NewTopPageHandler(topPageUseCase usecase.TopPageUseCase, c *request.Context) *topPageHandler {
	return &topPageHandler{
		topPageUseCase: topPageUseCase,
		C:              c,
	}
}

type responseTopPage struct {
	Article []model.Article `json:"article"`
	Pager   *pager.Pager    `json:"pager"`
}

func (tp topPageHandler) Get(w http.ResponseWriter, r *http.Request) error {
	currentPage := 1
	var res responseTopPage

	err := tp.C.Cache.GET(article_cache.TopPageAritcleListKey, &res)
	if err == nil {
		return tp.C.JSON(w, http.StatusOK, res)
	}

	// Number Of Articles Per 1 page
	limit := numberOfArticlePer1Page
	offset := numberOfArticlePer1Page * (currentPage - 1)
	err = tp.topPageUseCase.GetArticles(&res.Article, limit, offset)
	if err != nil {
		return tp.C.JSON(w, http.StatusInternalServerError, err.Error())
	}

	res.Pager, err = tp.topPageUseCase.GetPager(currentPage, numberOfArticlePer1Page)
	if err != nil {
		return tp.C.JSON(w, http.StatusInternalServerError, err.Error())
	}

	_ = tp.C.Cache.SET(article_cache.TopPageAritcleListKey, res)
	return tp.C.JSON(w, http.StatusOK, res)
}
