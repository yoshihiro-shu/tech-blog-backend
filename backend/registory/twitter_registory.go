package registory

import (
	"github.com/yoshihiro-shu/tech-blog-backend/backend/application/usecase"
	"github.com/yoshihiro-shu/tech-blog-backend/backend/infrastructure/persistence/postgres"
	"github.com/yoshihiro-shu/tech-blog-backend/backend/interfaces/api/handler"
	"github.com/yoshihiro-shu/tech-blog-backend/backend/interfaces/api/request"
)

func NewTwitterRegistory(ctx *request.Context) handler.TwitterHandler {
	twitterRepository := postgres.NewTwitterPersistence()
	twitterUseCase := usecase.NewTwitterUseCase(twitterRepository)
	return handler.NewTwitterHandler(twitterUseCase, ctx)
}
