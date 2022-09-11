package handler

import (
	"net/http"
	"strconv"

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
