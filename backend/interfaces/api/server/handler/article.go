package handler

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/yoshihiro-shu/draft-backend/backend/application/usecase"
	"github.com/yoshihiro-shu/draft-backend/backend/interfaces/api/server/request"
)

type ArticleHandler interface {
	Post(w http.ResponseWriter, r *http.Request) error
	Get(w http.ResponseWriter, r *http.Request) error
	Put(w http.ResponseWriter, r *http.Request) error
	Delete(w http.ResponseWriter, r *http.Request) error
}

type articleHandler struct {
	articleUseCase usecase.ArticleUseCase
	C              *request.Context
}

func NewArticleHandler(articleUseCase usecase.ArticleUseCase, c *request.Context) ArticleHandler {
	return &articleHandler{
		articleUseCase: articleUseCase,
		C:              c,
	}
}

func (ah *articleHandler) Post(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (ah *articleHandler) Get(w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)
	strId := vars["id"]
	id, err := strconv.Atoi(strId)
	if err != nil {
		return ah.C.JSON(w, http.StatusInternalServerError, err.Error())
	}

	article, err := ah.articleUseCase.FindByID(id)
	// TODO not no rows error
	if err != nil {
		return ah.C.JSON(w, http.StatusInternalServerError, err.Error())
	}

	return ah.C.JSON(w, http.StatusOK, article)
}

func (ah *articleHandler) Put(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (ah *articleHandler) Delete(w http.ResponseWriter, r *http.Request) error {
	return nil
}
