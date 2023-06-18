package registory

import (
	"github.com/go-pg/pg"
	"github.com/yoshihiro-shu/draft-backend/backend/application/usecase"
	"github.com/yoshihiro-shu/draft-backend/backend/infrastructure/persistence"
	"github.com/yoshihiro-shu/draft-backend/backend/interfaces/api/server/handler"
	"github.com/yoshihiro-shu/draft-backend/backend/interfaces/api/server/request"
	"github.com/yoshihiro-shu/draft-backend/backend/internal/pkg/logger"
	"gorm.io/gorm"
)

func NewLatestArticlesRegistory(ctx *request.Context, l logger.Logger, master, reprica func() *pg.DB, primary func() *gorm.DB) handler.LatestArticlesHandler {
	articleRepository := persistence.NewArticlePersistence(master, reprica, primary)
	articleUseCase := usecase.NewArticleUseCase(articleRepository)
	return handler.NewLatestArticlesHandler(articleUseCase, ctx, l)
}
