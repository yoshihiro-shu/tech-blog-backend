package handler

import (
	"net/http"

	"github.com/yoshihiro-shu/draft-backend/backend/application/usecase"
	"github.com/yoshihiro-shu/draft-backend/backend/interfaces/api/request"
	"github.com/yoshihiro-shu/draft-backend/backend/internal/auth"
	"go.uber.org/zap"
)

type UserHandler interface {
	SignUp(w http.ResponseWriter, r *http.Request) error
	Login(w http.ResponseWriter, r *http.Request) error
	RefreshToken(w http.ResponseWriter, r *http.Request) error
}

type userHandler struct {
	*request.Context
	userUseCase usecase.UserUseCase
}

func NewUserHandler(userUseCase usecase.UserUseCase, c *request.Context) *userHandler {
	return &userHandler{
		Context:     c,
		userUseCase: userUseCase,
	}
}

type signUpReq struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type signUpRes struct {
}

func (h *userHandler) SignUp(w http.ResponseWriter, r *http.Request) error {
	var req signUpReq
	err := h.MustBind(r, &req)
	if err != nil {
		// h.Logger.Error("error invalid request body.", zap.Error(err))zw
		return h.Error(w, http.StatusBadRequest, err)
	}

	hash, _ := auth.GenerateBcryptPassword(req.Password)
	err = h.userUseCase.Create(req.Name, hash, req.Email)
	if err != nil {
		return h.JSON(w, http.StatusInternalServerError, err.Error())
	}

	return h.JSON(w, http.StatusOK, &signUpRes{})
}

type loginReq struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type loginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func (h *userHandler) Login(w http.ResponseWriter, r *http.Request) error {
	var req loginReq
	err := h.MustBind(r, &req)
	if err != nil {
		h.Logger.Error("error invalid request body.", zap.Error(err))
		return h.Error(w, http.StatusBadRequest, err)
	}

	token, err := h.userUseCase.Login(req.Email, req.Password)
	if err != nil {
		switch err {
		case auth.ErrInvalidPassword:
			h.Logger.Error("Invalid Password.", zap.Error(err))
			return h.Error(w, http.StatusUnauthorized, err)
		}
		h.Logger.Error("server error.", zap.Error(err))
		return h.Error(w, http.StatusInternalServerError, err)
	}

	res := loginResponse{
		AccessToken:  token.AccessToken.JwtToken(),
		RefreshToken: token.RefreshToken.JwtToken(),
	}
	return h.JSON(w, http.StatusOK, res)
}

type refreshTokenReq struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}

type refreshTokenRes struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func (h *userHandler) RefreshToken(w http.ResponseWriter, r *http.Request) error {
	var req refreshTokenReq
	err := h.MustBind(r, &req)
	if err != nil {
		return h.Error(w, http.StatusInternalServerError, err)
	}

	authToken, err := h.userUseCase.RefreshToken(req.RefreshToken)
	if err != nil {
		return h.Error(w, http.StatusInternalServerError, err)
	}

	res := refreshTokenRes{
		AccessToken:  authToken.AccessToken.JwtToken(),
		RefreshToken: authToken.RefreshToken.JwtToken(),
	}

	return h.JSON(w, http.StatusOK, res)
}
