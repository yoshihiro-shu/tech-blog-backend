package handler

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/yoshihiro-shu/draft-backend/interfaces/api/server/auth"
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
