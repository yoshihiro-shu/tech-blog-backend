package registory

import (
	"github.com/yoshihiro-shu/tech-blog-backend/src/application/usecase"
	"github.com/yoshihiro-shu/tech-blog-backend/src/infrastructure/persistence/cache"
	"github.com/yoshihiro-shu/tech-blog-backend/src/infrastructure/persistence/postgres"
	"github.com/yoshihiro-shu/tech-blog-backend/src/interfaces/api/handler"
	"github.com/yoshihiro-shu/tech-blog-backend/src/internal/logger"
	"gorm.io/gorm"
)

func NewLatestArticlesRegistory(redis cache.RedisClient, l logger.Logger, master, reprica func() *gorm.DB) handler.LatestArticlesHandler {
	articleRepository := postgres.NewArticlePersistence(master, reprica)
	cacheArticleRepository := cache.NewArticleCacheAdaptor(redis)
	cacheArticlesRepository := cache.NewArticlesCacheAdaptor(redis)
	articleUseCase := usecase.NewArticleUseCase(articleRepository, cacheArticleRepository, cacheArticlesRepository)
	return handler.NewLatestArticlesHandler(articleUseCase, l)
}
