package registory

import (
	"github.com/yoshihiro-shu/draft-backend/backend/application/usecase"
	"github.com/yoshihiro-shu/draft-backend/backend/infrastructure/persistence/cache"
	"github.com/yoshihiro-shu/draft-backend/backend/infrastructure/persistence/postgres"
	"github.com/yoshihiro-shu/draft-backend/backend/interfaces/api/handler"
	"github.com/yoshihiro-shu/draft-backend/backend/interfaces/api/request"
	"github.com/yoshihiro-shu/draft-backend/backend/internal/logger"
	"gorm.io/gorm"
)

func NewTopPageRegistory(ctx *request.Context, l logger.Logger, master, reprica func() *gorm.DB) handler.TopPageHandler {
	articleRepository := postgres.NewArticlePersistence(master, reprica)
	cacheArticleRepository := cache.NewArticleCacheAdaptor(ctx.Cache())
	articleUseCase := usecase.NewArticleUseCase(articleRepository, cacheArticleRepository)
	return handler.NewTopPageHandler(articleUseCase, ctx, l)
}
