package registory

import (
	"github.com/yoshihiro-shu/tech-blog-backend/src/application/usecase"
	"github.com/yoshihiro-shu/tech-blog-backend/src/infrastructure/persistence/cache"
	"github.com/yoshihiro-shu/tech-blog-backend/src/infrastructure/persistence/postgres"
	"github.com/yoshihiro-shu/tech-blog-backend/src/interfaces/api/handler"
	"github.com/yoshihiro-shu/tech-blog-backend/src/interfaces/api/request"
	"gorm.io/gorm"
)

func NewArticleRegistory(ctx *request.Context, master, reprica func() *gorm.DB) handler.ArticleHandler {
	articleRepository := postgres.NewArticlePersistence(master, reprica)
	cacheArticleRepository := cache.NewArticleCacheAdaptor(ctx.Cache())
	cacheArticlesRepository := cache.NewArticlesCacheAdaptor(ctx.Cache())
	articleUseCase := usecase.NewArticleUseCase(articleRepository, cacheArticleRepository)
	articlesUseCase := usecase.NewArticlesUseCase(articleRepository, cacheArticlesRepository)
	return handler.NewArticleHandler(articleUseCase, articlesUseCase, ctx)
}
