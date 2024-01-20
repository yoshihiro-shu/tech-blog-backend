package handler

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/yoshihiro-shu/tech-blog-backend/src/application/usecase"
	"github.com/yoshihiro-shu/tech-blog-backend/src/domain/model"
	"github.com/yoshihiro-shu/tech-blog-backend/src/interfaces/api/request"
	"github.com/yoshihiro-shu/tech-blog-backend/src/internal/logger"
	"github.com/yoshihiro-shu/tech-blog-backend/src/internal/pager"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type LatestArticlesHandler interface {
	Get(w http.ResponseWriter, r *http.Request) error
}

type latestArticlesHandler struct {
	logger         logger.Logger
	articleUseCase usecase.ArticleUseCase
}

const (
	numberOfArticlePerPageAtLatestAritcles = 5
)

func NewLatestArticlesHandler(articleUseCase usecase.ArticleUseCase, l logger.Logger) LatestArticlesHandler {
	return &latestArticlesHandler{
		logger:         l,
		articleUseCase: articleUseCase,
	}
}

type responseLatestAritcles struct {
	Articles []model.Article `json:"articles"`
	Pager    *pager.Pager    `json:"pager"`
}

// LatestArticlesHandler godoc
// @Summary latest_articles handlers
// @Description get the latest articles by page
// @Accept  json
// @Produce  json
// @Success 200 {object} request.JSONResponce{data=responseLatestAritcles}
// @Failure 400 {object} request.JSONResponce{data=string}
// @Failure 500 {object} request.JSONResponce{data=string}
// @Router /api/new/{page} [get]
func (h latestArticlesHandler) Get(w http.ResponseWriter, r *http.Request) error {
	var res responseLatestAritcles
	vars := mux.Vars(r)
	strPage := vars["page"]
	currentPage, err := strconv.Atoi(strPage)
	if err != nil {
		h.logger.Error("failed at convert string to integer.", zap.Error(err))
		currentPage = 1
	}

	limit := numberOfArticlePerPageAtLatestAritcles
	offset := limit * (currentPage - 1)
	err = h.articleUseCase.GetArticles(&res.Articles, limit, offset, currentPage)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			h.logger.Warn("err no articles at latest Articles Handler")
			return request.JSON(w, http.StatusNotFound, err)
		}
		h.logger.Error("failed at get articles at latest articles.", zap.Error(err))
		return request.Error(w, http.StatusInternalServerError, err)
	}

	res.Pager, err = h.articleUseCase.GetPager(currentPage, limit)
	if err != nil {
		h.logger.Warn("failed at get pager at top page.", zap.Error(err))
		return request.Error(w, http.StatusInternalServerError, err)
	}

	return request.JSON(w, http.StatusOK, res)

}
