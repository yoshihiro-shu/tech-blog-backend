package registory

import (
	"github.com/go-pg/pg"
	"github.com/yoshihiro-shu/draft-backend/backend/application/usecase"
	"github.com/yoshihiro-shu/draft-backend/backend/infrastructure/persistence"
	"github.com/yoshihiro-shu/draft-backend/backend/interfaces/api/server/handler"
	"github.com/yoshihiro-shu/draft-backend/backend/interfaces/api/server/request"
	"github.com/yoshihiro-shu/draft-backend/backend/internal/pkg/logger"
)

func NewTopPageRegistory(ctx *request.Context, l logger.Logger, master, reprica func() *pg.DB) handler.TopPageHandler {
	topPageRepository := persistence.NewTopPagePersistence(master, reprica)
	topPageUseCase := usecase.NewTopPageUseCase(topPageRepository)
	return handler.NewTopPageHandler(topPageUseCase, ctx, l)
}
