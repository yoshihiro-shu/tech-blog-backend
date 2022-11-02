package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/yoshihiro-shu/draft-backend/application/usecase"
	"github.com/yoshihiro-shu/draft-backend/interfaces/api/server/auth"
	"github.com/yoshihiro-shu/draft-backend/interfaces/api/server/request"
	"github.com/yoshihiro-shu/draft-backend/internal/model"
	article_linkages_to_category "github.com/yoshihiro-shu/draft-backend/internal/model/article/linkages/to/category"
	article_linkages_to_many "github.com/yoshihiro-shu/draft-backend/internal/model/article/linkages/to/many"
)

func (h Handler) PostArticle(w http.ResponseWriter, r *http.Request) error {
	user_id := r.Context().Value(auth.UserKey)
	title := r.FormValue("title")
	content := r.FormValue("content")

	userId, err := strconv.Atoi(user_id.(string))
	if err != nil {
		return h.Context.JSON(w, http.StatusInternalServerError, err.Error())
	}

	a := model.Article{
		UserId:  userId,
		Title:   title,
		Content: content,
	}
	err = a.Insert(h.Context.Db.PsqlDB)
	if err != nil {
		return h.Context.JSON(w, http.StatusInternalServerError, err.Error())
	}

	return h.Context.JSON(w, http.StatusOK, a)
}

func (h Handler) GetArticles(w http.ResponseWriter, r *http.Request) error {
	a := new(article_linkages_to_category.Article)
	articles, err := a.GetList(h.Context.Db.PsqlDB)

	if err != nil {
		return h.Context.JSON(w, http.StatusInternalServerError, err.Error())
	}

	return h.Context.JSON(w, http.StatusOK, articles)
}

func (h Handler) GetArticleByID(w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)
	strId := vars["id"]
	id, _ := strconv.Atoi(strId)

	article := article_linkages_to_many.New(id)

	err := article.GetArticle(h.Context.Db.PsqlDB)
	if err != nil {
		return h.Context.JSON(w, http.StatusBadRequest, err.Error())
	}

	return h.Context.JSON(w, http.StatusOK, article)
}

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

type requestArticle struct {
}

type responseArticle struct {
}

func (ah articleHandler) Post(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (ah articleHandler) Get(w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)
	strId := vars["id"]
	id, err := strconv.Atoi(strId)
	fmt.Println("id", id)
	if err != nil {
		return ah.C.JSON(w, http.StatusInternalServerError, err.Error())
	}

	article, err := ah.articleUseCase.FindByID(id)
	if err != nil {
		return ah.C.JSON(w, http.StatusInternalServerError, err.Error())
	}

	return ah.C.JSON(w, http.StatusOK, article)
}

func (ah articleHandler) Put(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (ah articleHandler) Delete(w http.ResponseWriter, r *http.Request) error {
	return nil
}
