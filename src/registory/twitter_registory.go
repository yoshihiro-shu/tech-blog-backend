package registory

import (
	"github.com/yoshihiro-shu/tech-blog-backend/src/application/usecase"
	"github.com/yoshihiro-shu/tech-blog-backend/src/infrastructure/persistence/postgres"
	"github.com/yoshihiro-shu/tech-blog-backend/src/interfaces/api/handler"
	"github.com/yoshihiro-shu/tech-blog-backend/src/interfaces/api/request"
)

func NewTwitterRegistory(ctx *request.Context) handler.TwitterHandler {
	twitterRepository := postgres.NewTwitterPersistence()
	twitterUseCase := usecase.NewTwitterUseCase(twitterRepository)
	return handler.NewTwitterHandler(twitterUseCase, ctx)
}
