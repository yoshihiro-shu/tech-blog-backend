package registory

import (
	"github.com/yoshihiro-shu/tech-blog-backend/src/application/usecase"
	"github.com/yoshihiro-shu/tech-blog-backend/src/infrastructure/persistence/postgres"
	"github.com/yoshihiro-shu/tech-blog-backend/src/interfaces/api/handler"
	"github.com/yoshihiro-shu/tech-blog-backend/src/interfaces/api/request"
)

func NewUserRegistory(ctx *request.Context) handler.UserHandler {
	userRepository := postgres.NewUserPersistence(ctx.MasterDB, ctx.RepricaDB)
	refreshTokenRepository := postgres.NewRefreshTokenPersistence(ctx.MasterDB, ctx.RepricaDB)
	userUseCase := usecase.NewUserUseCase(userRepository, refreshTokenRepository)
	return handler.NewUserHandler(userUseCase, ctx)
}
