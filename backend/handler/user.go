package handler

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/yoshihiro-shu/draft-backend/auth"
	"github.com/yoshihiro-shu/draft-backend/controller"
	"github.com/yoshihiro-shu/draft-backend/model"
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

	user := controller.NewUser(email, password)

	err := user.Login(h.Context.Db.PsqlDB)
	if err != nil {
		return h.Context.JSON(w, http.StatusUnauthorized, err.Error())
	}

	// create TOKEN
	token := auth.CreateToken(strconv.Itoa(user.ID))

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
		ID: id,
	}

	err := user.GetByID(h.Context.Db.PsqlDB)
	if err != nil {
		return h.Context.JSON(w, http.StatusBadRequest, err.Error())
	}

	return h.Context.JSON(w, http.StatusOK, user)
}
