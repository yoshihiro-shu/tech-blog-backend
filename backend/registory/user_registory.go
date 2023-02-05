package registory

import (
	"github.com/yoshihiro-shu/draft-backend/application/usecase"
	"github.com/yoshihiro-shu/draft-backend/infrastructure/persistence"
	"github.com/yoshihiro-shu/draft-backend/interfaces/api/server/handler"
	"github.com/yoshihiro-shu/draft-backend/interfaces/api/server/request"
)

func NewUserRegistory(ctx *request.Context) handler.UserHandler {
	userRepository := persistence.NewUserPersistence(ctx.MasterDB())
	userUseCase := usecase.NewUserUseCase(userRepository)
	return handler.NewUserHandler(userUseCase, ctx)
}
