package registory

import (
	"github.com/yoshihiro-shu/draft-backend/application/usecase"
	"github.com/yoshihiro-shu/draft-backend/infrastructure/persistence"
	"github.com/yoshihiro-shu/draft-backend/interfaces/api/server/handler"
	"github.com/yoshihiro-shu/draft-backend/interfaces/api/server/request"
)

func NewArticleRegistory(ctx *request.Context) handler.ArticleHandler {
	articleRepository := persistence.NewArticlePersistence(ctx.MasterDB())
	articleUseCase := usecase.NewArticleUseCase(articleRepository)
	return handler.NewArticleHandler(articleUseCase, ctx)
}
