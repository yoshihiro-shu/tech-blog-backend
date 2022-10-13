package router

import (
	"net/http"

	"github.com/yoshihiro-shu/draft-backend/auth"
	"github.com/yoshihiro-shu/draft-backend/handler"
)

func (r Router) ApplyRouters() {
	r.Use(r.Context.TestMiddleware)

	h := handler.Handler{
		Context: r.Context,
	}

	r.AppHandle("/", h.Index).Methods(http.MethodGet)
	r.AppHandle("/article/{id}", h.GetArticleByID).Methods(http.MethodGet)

	// Grouping
	t := r.Group("/test")
	t.AppHandle("", h.TestHandler).Methods(http.MethodGet)
	t.AppHandle("/redis", h.TestSetRedis).Methods(http.MethodPost)
	t.AppHandle("/redis/{key}", h.TestGetRedis).Methods(http.MethodGet)
	t.AppHandle("/v2", h.Index).Methods(http.MethodGet)

	c := r.Group("/cmd")
	c.AppHandle("", h.Command).Methods(http.MethodGet)

	user := r.Group("/users")
	user.AppHandle("", h.GetUsers).Methods(http.MethodGet)
	user.AppHandle("/{id}", h.GetUserBYID).Methods(http.MethodGet)
	user.AppHandle("/login", h.Login).Methods(http.MethodPost)
	user.AppHandle("/signup", h.SignUp).Methods(http.MethodPost)
	// user.HandleFunc("/register", h.RegisterAccount).Methods(http.MethodPost)

	article := r.Group("/articles")
	article.Use(auth.AuthMiddleware)
	article.AppHandle("", h.PostArticle).Methods(http.MethodPost)
	article.AppHandle("", h.GetArticles).Methods(http.MethodGet)

	a := r.Group("/auth")
	a.Use(auth.AuthMiddleware)
	a.AppHandle("/index", h.AuthIndex).Methods(http.MethodGet)

}
