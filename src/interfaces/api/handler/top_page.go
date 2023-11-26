package handler

import (
	"net/http"

	"github.com/yoshihiro-shu/tech-blog-backend/src/application/usecase"
	"github.com/yoshihiro-shu/tech-blog-backend/src/domain/model"
	"github.com/yoshihiro-shu/tech-blog-backend/src/infrastructure/persistence/cache"
	"github.com/yoshihiro-shu/tech-blog-backend/src/interfaces/api/request"
	"github.com/yoshihiro-shu/tech-blog-backend/src/internal/logger"
	"go.uber.org/zap"
)

type TopPageHandler interface {
	Get(w http.ResponseWriter, r *http.Request) error
}

type topPageHandler struct {
	*request.Context
	logger         logger.Logger
	topPageUseCase usecase.TopPageUseCase
}

const (
	// 一ページあたりの記事数
	numberOfArticlePerPage = 10
)

func NewTopPageHandler(topPageUseCase usecase.TopPageUseCase, c *request.Context, l logger.Logger) TopPageHandler {
	return &topPageHandler{
		Context:        c,
		logger:         l,
		topPageUseCase: topPageUseCase,
	}
}

type responseTopPage struct {
	Article []model.Article `json:"articles"`
}

func (h topPageHandler) Get(w http.ResponseWriter, r *http.Request) error {
	currentPage := 1
	var res responseTopPage

	resKey := cache.TopPageKey()
	err := h.Cache().GET(resKey, &res)
	if err == nil {
		return h.JSON(w, http.StatusOK, res)
	}

	// Number Of Articles Per 1 page
	limit := numberOfArticlePerPage
	offset := numberOfArticlePerPage * (currentPage - 1)
	err = h.topPageUseCase.GetArticles(&res.Article, limit, offset)
	if err != nil {
		h.logger.Warn("failed at get articles at top page.", zap.Error(err))
		return h.JSON(w, http.StatusInternalServerError, err.Error())
	}

	err = h.Cache().SET(resKey, res)
	if err != nil {
		h.logger.Error("failed at set cache redis at top page.", zap.Error(err))
	}
	return h.JSON(w, http.StatusOK, res)
}
