package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/yoshihiro-shu/draft-backend/application/usecase"
	"github.com/yoshihiro-shu/draft-backend/controllers"
	"github.com/yoshihiro-shu/draft-backend/interfaces/api/server/auth"
	"github.com/yoshihiro-shu/draft-backend/interfaces/api/server/request"
	"github.com/yoshihiro-shu/draft-backend/internal/model"
)

type LoginResponse struct {
	Token string `json:"token"`
}

func (h Handler) SignUp(w http.ResponseWriter, r *http.Request) error {
	name := r.FormValue("name")
	email := r.FormValue("email")
	password := r.FormValue("password")

	user := model.User{
		Name:     name,
		Email:    email,
		Password: password,
	}
	user.SetBcryptPassword()

	err := user.Insert(h.Context.Db.PsqlDB)
	if err != nil {
		return h.Context.JSON(w, http.StatusInternalServerError, err.Error())
	}

	return h.Context.JSON(w, http.StatusOK, user)
}

func (h Handler) Login(w http.ResponseWriter, r *http.Request) error {
	email := r.FormValue("email")
	password := r.FormValue("password")

	user := controllers.NewUser(email, password)

	err := user.Login(h.Context.Db.PsqlDB)
	if err != nil {
		return h.Context.JSON(w, http.StatusUnauthorized, err.Error())
	}

	// create TOKEN
	token := auth.CreateToken(strconv.Itoa(user.Id))

	res := LoginResponse{Token: token}
	return h.Context.JSON(w, http.StatusOK, res)
}

func (h Handler) GetUsers(w http.ResponseWriter, r *http.Request) error {

	var u model.User
	users, err := u.GetAll(h.Context.Db.PsqlDB)
	if err != nil {
		return h.Context.JSON(w, http.StatusBadRequest, err.Error())
	}

	return h.Context.JSON(w, http.StatusOK, users)
}

func (h Handler) GetUserBYEmail(w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)
	email := vars["email"]

	user := &model.User{
		Email: email,
	}

	err := user.GetByEmail(h.Context.Db.PsqlDB)
	if err != nil {
		return h.Context.JSON(w, http.StatusBadRequest, err.Error())
	}

	return h.Context.JSON(w, http.StatusOK, user)
}

func (h Handler) GetUserBYID(w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)
	strId := vars["id"]
	id, _ := strconv.Atoi(strId)

	user := &model.User{
		Id: id,
	}

	err := user.GetByID(h.Context.Db.PsqlDB)
	if err != nil {
		return h.Context.JSON(w, http.StatusBadRequest, err.Error())
	}

	return h.Context.JSON(w, http.StatusOK, user)
}

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

func (uh *userHandler) SignUp(w http.ResponseWriter, r *http.Request) error {
	name := r.FormValue("name")
	email := r.FormValue("email")
	password := r.FormValue("password")

	hash, _ := auth.GenerateBcryptPassword(password)

	err := uh.userUseCase.Create(name, email, hash)
	if err != nil {
		return uh.C.JSON(w, http.StatusInternalServerError, err.Error())
	}

	return uh.C.JSON(w, http.StatusOK, nil)
}

func (uh *userHandler) Login(w http.ResponseWriter, r *http.Request) error {
	email := r.FormValue("email")
	password := r.FormValue("password")

	user, _ := uh.userUseCase.FindByEmail(email)

	fmt.Println(*user)

	// TODO fix here
	ok := auth.IsVerifyPassword(password, user.Password)
	if !ok {
		return uh.C.JSON(w, http.StatusUnauthorized, "your password is invalid")
	}

	// create TOKEN
	token := auth.CreateToken(strconv.Itoa(user.Id))

	res := LoginResponse{Token: token}
	return uh.C.JSON(w, http.StatusOK, res)
}
