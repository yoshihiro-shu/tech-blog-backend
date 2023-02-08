package registory

import (
	"github.com/go-pg/pg"
	"github.com/yoshihiro-shu/draft-backend/application/usecase"
	"github.com/yoshihiro-shu/draft-backend/infrastructure/persistence"
	"github.com/yoshihiro-shu/draft-backend/interfaces/api/server/handler"
	"github.com/yoshihiro-shu/draft-backend/interfaces/api/server/request"
)

func NewTopPageRegistory(ctx *request.Context, master, reprica func() *pg.DB) handler.TopPageHandler {
	topPageRepository := persistence.NewTopPagePersistence(master, reprica)
	topPageUseCase := usecase.NewTopPageUseCase(topPageRepository)
	return handler.NewTopPageHandler(topPageUseCase, ctx)
}
