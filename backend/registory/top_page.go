package registory

import (
	"github.com/yoshihiro-shu/draft-backend/application/usecase"
	"github.com/yoshihiro-shu/draft-backend/infrastructure/persistence"
	"github.com/yoshihiro-shu/draft-backend/interfaces/api/server/handler"
	"github.com/yoshihiro-shu/draft-backend/interfaces/api/server/request"
)

func NewTopPageRegistory(ctx *request.Context) handler.TopPageHandler {
	topPageRepository := persistence.NewTopPagePersistence(ctx.DB())
	topPageUseCase := usecase.NewTopPageUseCase(topPageRepository)
	return handler.NewTopPageHandler(topPageUseCase, ctx)
}
