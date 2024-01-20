package registory

import (
	"github.com/yoshihiro-shu/tech-blog-backend/src/application/usecase"
	"github.com/yoshihiro-shu/tech-blog-backend/src/infrastructure/persistence/cache"
	"github.com/yoshihiro-shu/tech-blog-backend/src/infrastructure/persistence/postgres"
	"github.com/yoshihiro-shu/tech-blog-backend/src/interfaces/api/handler"
	"github.com/yoshihiro-shu/tech-blog-backend/src/internal/logger"
	"gorm.io/gorm"
)

func NewArticleRegistory(redis cache.RedisClient, master, reprica func() *gorm.DB, logger logger.Logger) handler.ArticleHandler {
	articleRepository := postgres.NewArticlePersistence(master, reprica)
	cacheArticleRepository := cache.NewArticleCacheAdaptor(redis)
	cacheArticlesRepository := cache.NewArticlesCacheAdaptor(redis)
	articleUseCase := usecase.NewArticleUseCase(articleRepository, cacheArticleRepository, cacheArticlesRepository)
	articlesUseCase := usecase.NewArticlesUseCase(articleRepository, cacheArticlesRepository)
	return handler.NewArticleHandler(articleUseCase, articlesUseCase, logger)
}
