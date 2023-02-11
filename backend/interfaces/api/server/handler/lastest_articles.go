package handler

import (
	"net/http"
	"strconv"

	"github.com/go-pg/pg"
	"github.com/yoshihiro-shu/draft-backend/application/usecase"
	"github.com/yoshihiro-shu/draft-backend/domain/model"
	"github.com/yoshihiro-shu/draft-backend/interfaces/api/server/request"
	"github.com/yoshihiro-shu/draft-backend/internal/pkg/pager"
)

type LatestArticlesHandler interface {
	Get(w http.ResponseWriter, r *http.Request) error
}

type latestArticlesHandler struct {
	*request.Context
	articleUseCase usecase.ArticleUseCase
}

const (
	// TODO refactor
	numberOfArticlePerPageAtLatestAritcles = 1
)

func NewLatestArticlesHandler(articleUseCase usecase.ArticleUseCase, c *request.Context) LatestArticlesHandler {
	return &latestArticlesHandler{
		Context:        c,
		articleUseCase: articleUseCase,
	}
}

type responseLatestAritcles struct {
	Articles []model.Article `json:"articles"`
	Pager    *pager.Pager    `json:"pager"`
}

func (h latestArticlesHandler) Get(w http.ResponseWriter, r *http.Request) error {
	var res responseLatestAritcles

	page := r.URL.Query().Get("page")
	currentPage, err := strconv.Atoi(page)
	if err != nil {
		return h.Error(w, http.StatusBadRequest, err)
	}

	limit := numberOfArticlePerPageAtLatestAritcles
	offset := limit * (currentPage - 1)

	err = h.articleUseCase.GetArticles(&res.Articles, limit, offset)
	if err != nil {
		if err == pg.ErrNoRows {
			return h.JSON(w, http.StatusNotFound, err)
		}
		return h.Error(w, http.StatusInternalServerError, err)
	}

	res.Pager, err = h.articleUseCase.GetPager(currentPage, limit)
	if err != nil {
		return h.Error(w, http.StatusInternalServerError, err)
	}

	return h.JSON(w, http.StatusOK, res)

}
