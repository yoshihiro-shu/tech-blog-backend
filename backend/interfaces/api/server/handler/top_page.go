package handler

import (
	"fmt"
	"net/http"

	"github.com/yoshihiro-shu/draft-backend/application/usecase"
	"github.com/yoshihiro-shu/draft-backend/domain/model"
	"github.com/yoshihiro-shu/draft-backend/interfaces/api/server/request"
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
	articles *[]model.Article
}

func (tp topPageHandler) Get(w http.ResponseWriter, r *http.Request) error {
	var articles []model.Article

	err := tp.topPageUseCase.GetArticles(&articles)
	if err != nil {
		return tp.C.JSON(w, http.StatusInternalServerError, err.Error())
	}

	fmt.Println(articles)
	// res := &responseTopPage{
	// 	articles: &articles,
	// }

	return tp.C.JSON(w, http.StatusOK, articles)
}
