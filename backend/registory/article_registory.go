package registory

import (
	"github.com/yoshihiro-shu/draft-backend/backend/application/usecase"
	"github.com/yoshihiro-shu/draft-backend/backend/infrastructure/persistence"
	"github.com/yoshihiro-shu/draft-backend/backend/interfaces/api/handler"
	"github.com/yoshihiro-shu/draft-backend/backend/interfaces/api/request"
	"gorm.io/gorm"
)

func NewArticleRegistory(ctx *request.Context, master, reprica func() *gorm.DB) handler.ArticleHandler {
	articleRepository := persistence.NewArticlePersistence(master, reprica)
	articleUseCase := usecase.NewArticleUseCase(articleRepository, ctx.Cache())
	articlesUseCase := usecase.NewArticlesUseCase(articleRepository)
	return handler.NewArticleHandler(articleUseCase, articlesUseCase, ctx)
}
