package registory

import (
	"github.com/yoshihiro-shu/tech-blog-backend/backend/application/usecase"
	"github.com/yoshihiro-shu/tech-blog-backend/backend/infrastructure/persistence/postgres"
	"github.com/yoshihiro-shu/tech-blog-backend/backend/interfaces/api/handler"
	"github.com/yoshihiro-shu/tech-blog-backend/backend/interfaces/api/request"
)

func NewUserRegistory(ctx *request.Context) handler.UserHandler {
	userRepository := postgres.NewUserPersistence(ctx.MasterDB, ctx.RepricaDB)
	refreshTokenRepository := postgres.NewRefreshTokenPersistence(ctx.MasterDB, ctx.RepricaDB)
	userUseCase := usecase.NewUserUseCase(userRepository, refreshTokenRepository)
	return handler.NewUserHandler(userUseCase, ctx)
}
