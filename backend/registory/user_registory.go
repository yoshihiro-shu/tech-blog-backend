package registory

import (
	"github.com/yoshihiro-shu/draft-backend/backend/application/usecase"
	"github.com/yoshihiro-shu/draft-backend/backend/infrastructure/persistence"
	"github.com/yoshihiro-shu/draft-backend/backend/interfaces/api/server/handler"
	"github.com/yoshihiro-shu/draft-backend/backend/interfaces/api/server/request"
)

func NewUserRegistory(ctx *request.Context) handler.UserHandler {
	userRepository := persistence.NewUserPersistence(ctx.DBPrimary)
	refreshTokenRepository := persistence.NewRefreshTokenPersistence(ctx.MasterDB, ctx.RepricaDB, ctx.DBPrimary)
	userUseCase := usecase.NewUserUseCase(userRepository, refreshTokenRepository)
	return handler.NewUserHandler(userUseCase, ctx)
}
