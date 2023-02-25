package handler

import (
	"fmt"
	"net/http"

	"github.com/yoshihiro-shu/draft-backend/backend/application/usecase"
	"github.com/yoshihiro-shu/draft-backend/backend/domain/model"
	"github.com/yoshihiro-shu/draft-backend/backend/interfaces/api/server/cache"
	"github.com/yoshihiro-shu/draft-backend/backend/interfaces/api/server/request"
	"github.com/yoshihiro-shu/draft-backend/backend/internal/pkg/logger"
	"github.com/yoshihiro-shu/draft-backend/backend/internal/pkg/pager"
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
	Pager   *pager.Pager    `json:"pager"`
}

func (h topPageHandler) Get(w http.ResponseWriter, r *http.Request) error {
	currentPage := 1
	var res responseTopPage

	resKey := fmt.Sprintf(cache.TopPageAritcleListKeyByPage, currentPage)

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

	res.Pager, err = h.topPageUseCase.GetPager(currentPage, numberOfArticlePerPage)
	if err != nil {
		h.logger.Warn("failed at get pager at top page.", zap.Error(err))
		return h.JSON(w, http.StatusInternalServerError, err.Error())
	}

	err = h.Cache().SET(resKey, res)
	if err != nil {
		h.logger.Error("failed at set cache redis at top page.", zap.Error(err))
	}
	return h.JSON(w, http.StatusOK, res)
}
