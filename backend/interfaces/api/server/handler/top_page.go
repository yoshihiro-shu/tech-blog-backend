package handler

import (
	"fmt"
	"net/http"

	"github.com/yoshihiro-shu/draft-backend/application/usecase"
	"github.com/yoshihiro-shu/draft-backend/domain/model"
	"github.com/yoshihiro-shu/draft-backend/interfaces/api/server/cache"
	"github.com/yoshihiro-shu/draft-backend/interfaces/api/server/request"
	"github.com/yoshihiro-shu/draft-backend/internal/pkg/pager"
)

type TopPageHandler interface {
	Get(w http.ResponseWriter, r *http.Request) error
}

type topPageHandler struct {
	*request.Context
	topPageUseCase usecase.TopPageUseCase
}

const (
	// 一ページあたりの記事数
	numberOfArticlePerPage = 10
)

func NewTopPageHandler(topPageUseCase usecase.TopPageUseCase, c *request.Context) TopPageHandler {
	return &topPageHandler{
		Context:        c,
		topPageUseCase: topPageUseCase,
	}
}

type responseTopPage struct {
	Article []model.Article `json:"articles"`
	Pager   *pager.Pager    `json:"pager"`
}

func (tp topPageHandler) Get(w http.ResponseWriter, r *http.Request) error {
	currentPage := 1
	var res responseTopPage

	resKey := fmt.Sprintf(cache.TopPageAritcleListKeyByPage, currentPage)

	err := tp.Cache().GET(resKey, &res)
	if err == nil {
		return tp.JSON(w, http.StatusOK, res)
	}

	// Number Of Articles Per 1 page
	limit := numberOfArticlePerPage
	offset := numberOfArticlePerPage * (currentPage - 1)
	err = tp.topPageUseCase.GetArticles(&res.Article, limit, offset)
	if err != nil {
		return tp.JSON(w, http.StatusInternalServerError, err.Error())
	}

	res.Pager, err = tp.topPageUseCase.GetPager(currentPage, numberOfArticlePerPage)
	if err != nil {
		return tp.JSON(w, http.StatusInternalServerError, err.Error())
	}

	err = tp.Cache().SET(resKey, res)
	if err != nil {
		tp.Logger.Fatal(err.Error())
	}
	return tp.JSON(w, http.StatusOK, res)
}
