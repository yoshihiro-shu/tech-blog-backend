package registory

import (
	"github.com/yoshihiro-shu/draft-backend/backend/application/usecase"
	"github.com/yoshihiro-shu/draft-backend/backend/infrastructure/persistence"
	"github.com/yoshihiro-shu/draft-backend/backend/infrastructure/persistence/cache"
	"github.com/yoshihiro-shu/draft-backend/backend/interfaces/api/handler"
	"github.com/yoshihiro-shu/draft-backend/backend/interfaces/api/request"
	"gorm.io/gorm"
)

func NewArticleRegistory(ctx *request.Context, master, reprica func() *gorm.DB) handler.ArticleHandler {
	articleRepository := persistence.NewArticlePersistence(master, reprica)
	cacheArticleRepository := cache.NewArticleCacheAdaptor(ctx.Cache())
	cacheArticlesRepository := cache.NewArticlesCacheAdaptor(ctx.Cache())
	articleUseCase := usecase.NewArticleUseCase(articleRepository, cacheArticleRepository)
	articlesUseCase := usecase.NewArticlesUseCase(articleRepository, cacheArticlesRepository)
	return handler.NewArticleHandler(articleUseCase, articlesUseCase, ctx)
}
