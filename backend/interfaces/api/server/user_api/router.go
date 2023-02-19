package user_api

import (
	"github.com/yoshihiro-shu/draft-backend/interfaces/api/server/auth"
	"github.com/yoshihiro-shu/draft-backend/interfaces/api/server/cache"
	"github.com/yoshihiro-shu/draft-backend/interfaces/api/server/handler"
	"github.com/yoshihiro-shu/draft-backend/interfaces/api/server/middleware"
	"github.com/yoshihiro-shu/draft-backend/interfaces/api/server/model"
	"github.com/yoshihiro-shu/draft-backend/interfaces/api/server/request"
	"github.com/yoshihiro-shu/draft-backend/interfaces/api/server/router"
	"github.com/yoshihiro-shu/draft-backend/internal/config"
	"github.com/yoshihiro-shu/draft-backend/internal/pkg/logger"
	"github.com/yoshihiro-shu/draft-backend/registory"
)

func Apply(r router.Router, conf config.Configs, logger logger.Logger, db *model.DBContext, cache cache.RedisClient) {
	ctx := request.NewContext(conf, logger, db, cache)

	r.Use(middleware.CorsMiddleware)
	r.Use(middleware.LoggerMiddleware(logger))

	h := handler.NewIndexHandler(ctx)

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
		// Grouping
		t := r.Group("/test")
		t.GET("", h.TestHandler)
		t.POST("/redis", h.TestSetRedis)
		t.GET("/redis/{key}", h.TestGetRedis)
		t.GET("/v2", h.Index)
	}
	{
		c := r.Group("/cmd")
		c.GET("", h.Command)
	}
	{
		user := r.Group("/users")
		userHandler := registory.NewUserRegistory(ctx)
		user.POST("/login", userHandler.Login)
		user.POST("/signup", userHandler.SignUp)
		// user.POST("/register", h.RegisterAccount)
	}
	{
		articleHandler := registory.NewArticleRegistory(ctx, ctx.MasterDB, ctx.RepricaDB)
		article := r.Group("/articles")
		article.GET("/{id:[0-9]+}", articleHandler.Get)
	}
	{
		a := r.Group("/auth")
		a.Use(auth.AuthMiddleware)
		a.GET("/index", h.AuthIndex)
	}
}