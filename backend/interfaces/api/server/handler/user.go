package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-pg/pg"
	"github.com/yoshihiro-shu/draft-backend/application/usecase"
	"github.com/yoshihiro-shu/draft-backend/interfaces/api/server/auth"
	"github.com/yoshihiro-shu/draft-backend/interfaces/api/server/request"
)

type UserHandler interface {
	SignUp(w http.ResponseWriter, r *http.Request) error
	Login(w http.ResponseWriter, r *http.Request) error
}

type userHandler struct {
	userUseCase usecase.UserUseCase
	C           *request.Context
}

func NewUserHandler(userUseCase usecase.UserUseCase, c *request.Context) *userHandler {
	return &userHandler{
		userUseCase: userUseCase,
		C:           c,
	}
}

type requestUser struct {
}

type responseUser struct {
}

type loginResponse struct {
	Token string `json:"token"`
}

func (uh *userHandler) SignUp(w http.ResponseWriter, r *http.Request) error {
	name := r.FormValue("name")
	email := r.FormValue("email")
	password := r.FormValue("password")

	hash, _ := auth.GenerateBcryptPassword(password)

	err := uh.userUseCase.Create(name, hash, email)
	if err != nil {
		return uh.C.JSON(w, http.StatusInternalServerError, err.Error())
	}

	return uh.C.JSON(w, http.StatusOK, nil)
}

func (uh *userHandler) Login(w http.ResponseWriter, r *http.Request) error {
	email := r.FormValue("email")
	password := r.FormValue("password")

	user, err := uh.userUseCase.FindByEmail(email)
	if err != nil {
		if err == pg.ErrNoRows {
			return uh.C.JSON(w, http.StatusNotFound, err.Error())
		}
		return uh.C.JSON(w, http.StatusInternalServerError, err.Error())
	}

	fmt.Println(*user)

	// TODO fix here
	err = auth.IsVerifyPassword(password, user.Password)
	if err != nil {
		return uh.C.JSON(w, http.StatusUnauthorized, "your password is invalid")
	}

	// create TOKEN
	token := auth.CreateToken(strconv.Itoa(user.Id))

	res := loginResponse{Token: token}
	return uh.C.JSON(w, http.StatusOK, res)
}
