package handler

import (
	"net/http"

	"github.com/yoshihiro-shu/tech-blog-backend/src/application/usecase"
	"github.com/yoshihiro-shu/tech-blog-backend/src/domain/model"
	"github.com/yoshihiro-shu/tech-blog-backend/src/interfaces/api/request"
	"github.com/yoshihiro-shu/tech-blog-backend/src/internal/logger"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type TopPageHandler interface {
	Get(w http.ResponseWriter, r *http.Request) error
}

type topPageHandler struct {
	logger         logger.Logger
	articleUseCase usecase.ArticleUseCase
}

const (
	// 一ページあたりの記事数
	numberOfArticlePerPage = 10
)

func NewTopPageHandler(articleUseCase usecase.ArticleUseCase, l logger.Logger) LatestArticlesHandler {
	return &topPageHandler{
		logger:         l,
		articleUseCase: articleUseCase,
	}
}

type responseTopPage struct {
	Article []model.Article `json:"articles"`
}

// TopPageHandler godoc
// @Summary top_page handlers
// @Description get the top page articles
// @Accept  json
// @Produce  json
// @Success 200 {object} request.JSONResponce{data=responseTopPage}
// @Failure 400 {object} request.JSONResponce{data=string}
// @Failure 500 {object} request.JSONResponce{data=string}
// @Router /api/top_page [get]
// func (h topPageHandler) Get(w http.ResponseWriter, r *http.Request) error {
// 	currentPage := 1
// 	var res responseTopPage

// 	resKey := cache.TopPageKey()
// 	err := h.Cache().GET(resKey, &res)
// 	if err == nil {
// 		return h.JSON(w, http.StatusOK, res)
// 	}

// 	// Number Of Articles Per 1 page
// 	limit := numberOfArticlePerPage
// 	offset := numberOfArticlePerPage * (currentPage - 1)
// 	err = h.topPageUseCase.GetArticles(&res.Article, limit, offset)
// 	if err != nil {
// 		h.logger.Warn("failed at get articles at top page.", zap.Error(err))
// 		return h.JSON(w, http.StatusInternalServerError, err.Error())
// 	}

// 	err = h.Cache().SET(resKey, res)
// 	if err != nil {
// 		h.logger.Error("failed at set cache redis at top page.", zap.Error(err))
// 	}
// 	return h.JSON(w, http.StatusOK, res)
// }

func (h topPageHandler) Get(w http.ResponseWriter, r *http.Request) error {
	currentPage := 1
	var res responseTopPage

	// Number Of Articles Per 1 page
	limit := numberOfArticlePerPage
	offset := numberOfArticlePerPage * (currentPage - 1)
	err := h.articleUseCase.GetArticles(&res.Article, limit, offset, currentPage)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			h.logger.Warn("err no articles at latest Articles Handler")
			return request.JSON(w, http.StatusNotFound, err)
		}
		h.logger.Error("failed at get articles at latest articles.", zap.Error(err))
		return request.Error(w, http.StatusInternalServerError, err)
	}

	return request.JSON(w, http.StatusOK, res)
}
