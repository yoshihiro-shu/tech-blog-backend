package user_api

import (
	"github.com/yoshihiro-shu/draft-backend/backend/infrastructure/persistence/cache"
	"github.com/yoshihiro-shu/draft-backend/backend/interfaces/api/handler"
	"github.com/yoshihiro-shu/draft-backend/backend/interfaces/api/middlewares"
	"github.com/yoshihiro-shu/draft-backend/backend/interfaces/api/request"
	"github.com/yoshihiro-shu/draft-backend/backend/interfaces/api/router"
	"github.com/yoshihiro-shu/draft-backend/backend/interfaces/model"
	"github.com/yoshihiro-shu/draft-backend/backend/internal/config"
	"github.com/yoshihiro-shu/draft-backend/backend/internal/logger"
	"github.com/yoshihiro-shu/draft-backend/backend/registory"
)

func Apply(r router.Router, conf config.Configs, logger logger.Logger, db model.DBClient, cache cache.RedisClient) {
	ctx := request.NewContext(conf, logger, db, cache)

	r.Use(middlewares.Logger(logger))
	r.Use(middlewares.CsrfProtecter(conf, logger))
	r.Use(middlewares.SetterCsrfToken)
	r.Use(middlewares.Cors(conf.Frontend))

	h := handler.NewIndexHandler(ctx)

	r = r.Group("/api")
	{
		r.GET("/healthcheck", h.Index)
	}
	{
		topPageHandler := registory.NewTopPageRegistory(
			ctx,
			logger,
			ctx.MasterDB,
			ctx.RepricaDB,
		)
		r.GET("/top", topPageHandler.Get)
	}
	{
		lastestAriclesHandler := registory.NewLatestArticlesRegistory(
			ctx,
			logger,
			ctx.MasterDB,
			ctx.RepricaDB,
		)
		latestArticles := r.Group("/new")
		latestArticles.GET("/{page:[0-9]+}", lastestAriclesHandler.Get)
	}
	{
		twitterHandler := registory.NewTwitterRegistory(ctx)
		twitter := r.Group("/twitter")
		twitter.GET("/timeline", twitterHandler.GetTimeLine)
	}
	{
		auth := r.Group("/auth")
		userHandler := registory.NewUserRegistory(ctx)
		auth.POST("/login", userHandler.Login)
		auth.POST("/signup", userHandler.SignUp)
		auth.POST("/refresh_token", userHandler.RefreshToken)
	}
	{
		articleHandler := registory.NewArticleRegistory(
			ctx,
			ctx.MasterDB,
			ctx.RepricaDB,
		)
		article := r.Group("/articles")
		article.GET("/{id:[0-9]+}", articleHandler.Get)
		// article.GET("/category/{slug}/{id:[0-9]+}", articleHandler.GetArticlesByCategory)
		// article.GET("/tag/{slug}/{id:[0-9]+}", articleHandler.GetArticlesByTag)
		article.GET("/category/{slug}", articleHandler.GetArticlesByCategory)
		article.GET("/tag/{slug}", articleHandler.GetArticlesByTag)
		// article.GET("/category/{slug}/{id:[0-9]+}", articleHandler.GetArticlesByCategory)
		// article.GET("/tag/{slug}/{id:[0-9]+}", articleHandler.GetArticlesByTag)
	}
	// {
	// 	a := r.Group("/auth")
	// 	a.Use(middleware.Auth)
	// 	a.GET("/index", h.AuthIndex)
	// }
}
