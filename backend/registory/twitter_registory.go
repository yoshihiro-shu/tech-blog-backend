package registory

import (
	"github.com/yoshihiro-shu/draft-backend/application/usecase"
	"github.com/yoshihiro-shu/draft-backend/infrastructure/persistence"
	"github.com/yoshihiro-shu/draft-backend/interfaces/api/server/handler"
	"github.com/yoshihiro-shu/draft-backend/interfaces/api/server/request"
)

func NewTwitterRegistory(ctx *request.Context) handler.TwitterHandler {
	twitterRepository := persistence.NewTwitterPersistence()
	twitterUseCase := usecase.NewTwitterUseCase(twitterRepository)
	return handler.NewTwitterHandler(twitterUseCase, ctx)
}
