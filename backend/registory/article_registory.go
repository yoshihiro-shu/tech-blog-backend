package registory

import (
	"github.com/go-pg/pg"
	"github.com/yoshihiro-shu/draft-backend/backend/application/usecase"
	"github.com/yoshihiro-shu/draft-backend/backend/infrastructure/persistence"
	"github.com/yoshihiro-shu/draft-backend/backend/interfaces/api/server/handler"
	"github.com/yoshihiro-shu/draft-backend/backend/interfaces/api/server/request"
	"gorm.io/gorm"
)

func NewArticleRegistory(ctx *request.Context, master, reprica func() *pg.DB, primary func() *gorm.DB) handler.ArticleHandler {
	articleRepository := persistence.NewArticlePersistence(master, reprica, primary)
	articleUseCase := usecase.NewArticleUseCase(articleRepository)
	return handler.NewArticleHandler(articleUseCase, ctx)
}
