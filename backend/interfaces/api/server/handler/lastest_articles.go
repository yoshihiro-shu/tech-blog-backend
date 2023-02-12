package handler

import (
	"net/http"
	"strconv"

	"github.com/go-pg/pg"
	"github.com/yoshihiro-shu/draft-backend/application/usecase"
	"github.com/yoshihiro-shu/draft-backend/domain/model"
	"github.com/yoshihiro-shu/draft-backend/interfaces/api/server/request"
	"github.com/yoshihiro-shu/draft-backend/internal/pkg/logger"
	"github.com/yoshihiro-shu/draft-backend/internal/pkg/pager"
	"go.uber.org/zap"
)

type LatestArticlesHandler interface {
	Get(w http.ResponseWriter, r *http.Request) error
}

type latestArticlesHandler struct {
	*request.Context
	logger         logger.Logger
	articleUseCase usecase.ArticleUseCase
}

const (
	// TODO refactor
	numberOfArticlePerPageAtLatestAritcles = 1
)

func NewLatestArticlesHandler(articleUseCase usecase.ArticleUseCase, c *request.Context, l logger.Logger) LatestArticlesHandler {
	return &latestArticlesHandler{
		Context:        c,
		logger:         l,
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
		h.logger.Error("failed at convert string to integer.", zap.Error(err))
		currentPage = 1
	}

	limit := numberOfArticlePerPageAtLatestAritcles
	offset := limit * (currentPage - 1)

	err = h.articleUseCase.GetArticles(&res.Articles, limit, offset)
	if err != nil {
		if err == pg.ErrNoRows {
			h.logger.Warn("err no articles at latest Articles Handler")
			return h.JSON(w, http.StatusNotFound, err)
		}
		h.logger.Warn("failed at get articles at latest articles.", zap.Error(err))
		return h.Error(w, http.StatusInternalServerError, err)
	}

	res.Pager, err = h.articleUseCase.GetPager(currentPage, limit)
	if err != nil {
		h.logger.Warn("failed at get pager at top page.", zap.Error(err))
		return h.Error(w, http.StatusInternalServerError, err)
	}

	return h.JSON(w, http.StatusOK, res)

}
