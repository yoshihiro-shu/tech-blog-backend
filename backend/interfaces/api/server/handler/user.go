package handler

import (
	"net/http"

	"github.com/yoshihiro-shu/draft-backend/application/usecase"
	"github.com/yoshihiro-shu/draft-backend/interfaces/api/server/auth"
	"github.com/yoshihiro-shu/draft-backend/interfaces/api/server/request"
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

type requestUser struct {
}

type responseUser struct {
}

func (uh *userHandler) SignUp(w http.ResponseWriter, r *http.Request) error {
	name := r.FormValue("name")
	email := r.FormValue("email")
	password := r.FormValue("password")

	hash, _ := auth.GenerateBcryptPassword(password)

	err := uh.userUseCase.Create(name, hash, email)
	if err != nil {
		return uh.JSON(w, http.StatusInternalServerError, err.Error())
	}

	return uh.JSON(w, http.StatusOK, nil)
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
		return h.Error(w, http.StatusBadRequest, err)
	}

	token, err := h.userUseCase.Login(req.Email, req.Password)
	if err != nil {
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
