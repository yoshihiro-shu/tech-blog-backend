package registory

import (
	"github.com/yoshihiro-shu/tech-blog-backend/src/application/usecase"
	"github.com/yoshihiro-shu/tech-blog-backend/src/infrastructure/persistence/cache"
	"github.com/yoshihiro-shu/tech-blog-backend/src/infrastructure/persistence/postgres"
	"github.com/yoshihiro-shu/tech-blog-backend/src/interfaces/api/handler"
	"github.com/yoshihiro-shu/tech-blog-backend/src/interfaces/api/request"
	"github.com/yoshihiro-shu/tech-blog-backend/src/internal/logger"
	"gorm.io/gorm"
)

func NewTopPageRegistory(ctx *request.Context, l logger.Logger, master, reprica func() *gorm.DB) handler.TopPageHandler {
	articleRepository := postgres.NewArticlePersistence(master, reprica)
	cacheArticleRepository := cache.NewArticleCacheAdaptor(ctx.Cache())
	articleUseCase := usecase.NewArticleUseCase(articleRepository, cacheArticleRepository)
	return handler.NewTopPageHandler(articleUseCase, ctx, l)
}
