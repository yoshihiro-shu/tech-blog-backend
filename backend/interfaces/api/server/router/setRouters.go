package router

import (
	"net/http"

	"github.com/yoshihiro-shu/draft-backend/interfaces/api/server/auth"
	"github.com/yoshihiro-shu/draft-backend/interfaces/api/server/handler"
	"github.com/yoshihiro-shu/draft-backend/interfaces/api/server/middleware"
	"github.com/yoshihiro-shu/draft-backend/interfaces/api/server/request"
	"github.com/yoshihiro-shu/draft-backend/internal/config"
	"github.com/yoshihiro-shu/draft-backend/internal/pkg/logger"
	"github.com/yoshihiro-shu/draft-backend/registory"
)

func (r Router) Apply(conf config.Configs, logger logger.Logger) {
	ctx := request.NewContext(conf, logger)

	r.Use(middleware.CorsMiddleware)
	r.Use(middleware.LoggerMiddleware(logger))

	h := handler.NewIndexHandler(ctx)

	{
		r.AppHandle("/healthcheck", h.Index).Methods(http.MethodGet)
	}
	{
		topPageHandler := registory.NewTopPageRegistory(
			ctx,
			logger,
			ctx.MasterDB,
			ctx.RepricaDB,
		)
		r.AppHandle("/top", topPageHandler.Get).Methods(http.MethodGet)
	}
	{
		lastestAriclesHandler := registory.NewLatestArticlesRegistory(
			ctx,
			logger,
			ctx.MasterDB,
			ctx.RepricaDB,
		)
		latestArticles := r.Group("/new")
		latestArticles.AppHandle("", lastestAriclesHandler.Get).Methods(http.MethodGet)
	}
	{
		twitterHandler := registory.NewTwitterRegistory(ctx)
		twitter := r.Group("/twitter")
		twitter.AppHandle("/timeline", twitterHandler.GetTimeLine).Methods(http.MethodGet)
	}
	{
		// Grouping
		t := r.Group("/test")
		t.AppHandle("", h.TestHandler).Methods(http.MethodGet)
		t.AppHandle("/redis", h.TestSetRedis).Methods(http.MethodPost)
		t.AppHandle("/redis/{key}", h.TestGetRedis).Methods(http.MethodGet)
		t.AppHandle("/v2", h.Index).Methods(http.MethodGet)
	}
	{
		c := r.Group("/cmd")
		c.AppHandle("", h.Command).Methods(http.MethodGet)
	}
	{
		user := r.Group("/users")
		userHandler := registory.NewUserRegistory(ctx)
		user.AppHandle("/login", userHandler.Login).Methods(http.MethodPost)
		user.AppHandle("/signup", userHandler.SignUp).Methods(http.MethodPost)
		// user.HandleFunc("/register", h.RegisterAccount).Methods(http.MethodPost)
	}
	{
		articleHandler := registory.NewArticleRegistory(ctx, ctx.MasterDB, ctx.RepricaDB)
		article := r.Group("/articles")
		article.AppHandle("/{id:[0-9]+}", articleHandler.Get).Methods(http.MethodGet)
	}
	{
		a := r.Group("/auth")
		a.Use(auth.AuthMiddleware)
		a.AppHandle("/index", h.AuthIndex).Methods(http.MethodGet)
	}
}
