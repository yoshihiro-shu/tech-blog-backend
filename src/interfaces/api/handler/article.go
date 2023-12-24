package handler

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/yoshihiro-shu/tech-blog-backend/src/application/usecase"
	"github.com/yoshihiro-shu/tech-blog-backend/src/domain/model"
	"github.com/yoshihiro-shu/tech-blog-backend/src/interfaces/api/request"
	"gorm.io/gorm"
)

type ArticleHandler interface {
	Get(w http.ResponseWriter, r *http.Request) error
	GetArticlesByCategory(w http.ResponseWriter, r *http.Request) error
	GetArticlesByTag(w http.ResponseWriter, r *http.Request) error
}

type articleHandler struct {
	articleUseCase  usecase.ArticleUseCase
	articlesUseCase usecase.ArticlesUseCase
	C               *request.Context
}

// ArticleHandler godoc
// @Summary article handlers
// @Description get the article by id
// @Accept  json
// @Produce  json
// @Success 200 {object} request.JSONResponce{data=model.Article}
// @Failure 400 {object} request.JSONResponce{data=string}
// @Failure 500 {object} request.JSONResponce{data=string}
// @Router /api/articles/{id} [get]
func (ah *articleHandler) Get(w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)
	strId := vars["id"]
	id, err := strconv.Atoi(strId)
	if err != nil {
		return ah.C.JSON(w, http.StatusInternalServerError, err.Error())
	}

	article, err := ah.articleUseCase.FindByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return ah.C.JSON(w, http.StatusNotFound, err.Error())
		}
		return ah.C.JSON(w, http.StatusInternalServerError, err.Error())
	}

	return ah.C.JSON(w, http.StatusOK, article)
}

type responseGetArticlesByCategory struct {
	Articles []model.Article `json:"articles"`
}

// ArticleHandler godoc
// @Summary article handlers
// @Description get the article by category
// @Accept  json
// @Produce  json
// @Success 200 {object} request.JSONResponce{data=responseGetArticlesByCategory}
// @Failure 400 {object} request.JSONResponce{data=string}
// @Failure 500 {object} request.JSONResponce{data=string}
// @Router /api/articles/category/{slug} [get]
func (ah *articleHandler) GetArticlesByCategory(w http.ResponseWriter, r *http.Request) error {
	var res responseGetArticlesByCategory
	vars := mux.Vars(r)
	slug := vars["slug"]

	err := ah.articlesUseCase.GetArticlesByCategory(&res.Articles, slug)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			ah.C.Logger.Warn("err no articles at latest Articles Handler")
			return ah.C.JSON(w, http.StatusNotFound, err)
		}
	}

	return ah.C.JSON(w, http.StatusOK, res)
}

type responseGetArticlesByTag struct {
	Articles []model.Article `json:"articles"`
}

// ArticleHandler godoc
// @Summary article handlers
// @Description get the article by tag
// @Accept  json
// @Produce  json
// @Success 200 {object} request.JSONResponce{data=responseGetArticlesByTag}
// @Failure 400 {object} request.JSONResponce{data=string}
// @Failure 500 {object} request.JSONResponce{data=string}
// @Router /api/articles//{slug} [get]
func (ah *articleHandler) GetArticlesByTag(w http.ResponseWriter, r *http.Request) error {
	var res responseGetArticlesByTag
	vars := mux.Vars(r)
	slug := vars["slug"]

	err := ah.articlesUseCase.GetArticlesByTag(&res.Articles, slug)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			ah.C.Logger.Warn("err no articles at latest Articles Handler")
			return ah.C.JSON(w, http.StatusNotFound, err)
		}
	}

	return ah.C.JSON(w, http.StatusOK, res)
}

func NewArticleHandler(articleUseCase usecase.ArticleUseCase, articlesUseCase usecase.ArticlesUseCase, c *request.Context) ArticleHandler {
	return &articleHandler{
		articleUseCase:  articleUseCase,
		articlesUseCase: articlesUseCase,
		C:               c,
	}
}
