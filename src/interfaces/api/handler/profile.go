package handler

import (
	"net/http"

	"github.com/yoshihiro-shu/tech-blog-backend/src/application/usecase"
	"github.com/yoshihiro-shu/tech-blog-backend/src/interfaces/api/request"
	"github.com/yoshihiro-shu/tech-blog-backend/src/internal/logger"
	"go.uber.org/zap"
)

type ProfileHandler interface {
	GetResume(w http.ResponseWriter, r *http.Request) error
}

type profileHandler struct {
	logger         logger.Logger
	profileUsecase usecase.ProfileUseCase
}

func NewProfileHandler(profileUseCase usecase.ProfileUseCase, logger logger.Logger) ProfileHandler {
	return &profileHandler{
		profileUsecase: profileUseCase,
		logger:         logger,
	}
}

type responseGetResume struct {
	HTMLContent string `json:"htmlContent"`
}

func (h profileHandler) GetResume(w http.ResponseWriter, r *http.Request) error {
	res, err := h.profileUsecase.GetResume()
	if err != nil {
		return request.JSON(w, http.StatusInternalServerError, err.Error())
	}
	h.logger.Info("Get Resume String", zap.String("htmlContent", string(res)))
	h.logger.Info("Get Resume Any", zap.Any("res", res))
	return request.JSON(w, http.StatusOK, responseGetResume{string(res)})
}
