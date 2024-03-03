package registory

import (
	"github.com/yoshihiro-shu/tech-blog-backend/src/application/usecase"
	"github.com/yoshihiro-shu/tech-blog-backend/src/infrastructure/persistence/cache"
	"github.com/yoshihiro-shu/tech-blog-backend/src/interfaces/api/handler"
	"github.com/yoshihiro-shu/tech-blog-backend/src/internal/logger"
)

func NewProfileRegistory(token string, redis cache.RedisClient, logger logger.Logger) handler.ProfileHandler {
	profileCacheRepo := cache.NewProfileCacheAdaptor(redis)
	profileUsecase := usecase.NewProfileUseCase(token, profileCacheRepo)
	return handler.NewProfileHandler(profileUsecase, logger)
}
