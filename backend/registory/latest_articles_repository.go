package registory

import (
	"github.com/go-pg/pg"
	"github.com/yoshihiro-shu/draft-backend/application/usecase"
	"github.com/yoshihiro-shu/draft-backend/infrastructure/persistence"
	"github.com/yoshihiro-shu/draft-backend/interfaces/api/server/handler"
	"github.com/yoshihiro-shu/draft-backend/interfaces/api/server/request"
	"github.com/yoshihiro-shu/draft-backend/internal/pkg/logger"
)

func NewLatestArticlesRegistory(ctx *request.Context, l logger.Logger, master, reprica func() *pg.DB) handler.LatestArticlesHandler {
	articleRepository := persistence.NewArticlePersistence(master, reprica)
	articleUseCase := usecase.NewArticleUseCase(articleRepository)
	return handler.NewLatestArticlesHandler(articleUseCase, ctx, l)
}
