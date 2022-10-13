package handler

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/yoshihiro-shu/draft-backend/auth"
	"github.com/yoshihiro-shu/draft-backend/model"
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
	var a model.Article
	articles, err := a.GetAll(h.Context.Db.PsqlDB)
	if err != nil {
		return h.Context.JSON(w, http.StatusInternalServerError, err.Error())
	}
	return h.Context.JSON(w, http.StatusOK, articles)
}

func (h Handler) GetArticleByID(w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)
	strId := vars["id"]
	id, _ := strconv.Atoi(strId)

	article := &model.Article{
		Id: id,
	}

	err := article.GetByID(h.Context.Db.PsqlDB)
	if err != nil {
		return h.Context.JSON(w, http.StatusBadRequest, err.Error())
	}

	return h.Context.JSON(w, http.StatusOK, article)
}
