package handler

import (
	"net/http"

	"github.com/yoshihiro-shu/tech-blog-backend/src/application/usecase"
	"github.com/yoshihiro-shu/tech-blog-backend/src/domain/model"
	"github.com/yoshihiro-shu/tech-blog-backend/src/interfaces/api/request"
)

type TwitterHandler interface {
	GetTimeLine(w http.ResponseWriter, r *http.Request) error
}

type twitterHandler struct {
	twitterUseCase usecase.TwitterUseCase
	C              *request.Context
}

type resTweets struct {
	Tweets []model.Tweet `json:"tweets"`
}

func NewTwitterHandler(twitterUseCase usecase.TwitterUseCase, c *request.Context) TwitterHandler {
	return &twitterHandler{
		twitterUseCase: twitterUseCase,
		C:              c,
	}
}

func (th twitterHandler) GetTimeLine(w http.ResponseWriter, r *http.Request) error {
	timelines, err := th.twitterUseCase.GetTimelines(th.C.Conf)
	if err != nil {
		return th.C.JSON(w, http.StatusInternalServerError, err.Error())
	}

	res := resTweets{
		Tweets: timelines,
	}

	return th.C.JSON(w, http.StatusOK, res)
}
