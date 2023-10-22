package handler

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/yoshihiro-shu/draft-backend/backend/application/usecase"
	"github.com/yoshihiro-shu/draft-backend/backend/domain/model"
	"github.com/yoshihiro-shu/draft-backend/backend/interfaces/api/request"
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

type ListAritcleBox struct {
	Id           int       `json:"id"`
	Title        string    `json:"title"`
	ThumbnailUrl string    `json:"thumbnail_url"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	// CategoryName string    `json:"category_name"`
	// CategorySlug string    `json:"category_slug"`
}

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
