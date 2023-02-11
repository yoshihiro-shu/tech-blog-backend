package registory

import (
	"github.com/go-pg/pg"
	"github.com/yoshihiro-shu/draft-backend/application/usecase"
	"github.com/yoshihiro-shu/draft-backend/infrastructure/persistence"
	"github.com/yoshihiro-shu/draft-backend/interfaces/api/server/handler"
	"github.com/yoshihiro-shu/draft-backend/interfaces/api/server/request"
)

func NewArticleRegistory(ctx *request.Context, master, reprica func() *pg.DB) handler.ArticleHandler {
	articleRepository := persistence.NewArticlePersistence(master, reprica)
	articleUseCase := usecase.NewArticleUseCase(articleRepository)
	return handler.NewArticleHandler(articleUseCase, ctx)
}
