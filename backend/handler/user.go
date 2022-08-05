package handler

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/yoshi429/test/auth"
	"github.com/yoshi429/test/model"
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

	err := user.Insert(h.Context.Db.PSQLDB)
	if err != nil {
		return h.Context.JSON(w, http.StatusInternalServerError, err.Error())
	}

	return h.Context.JSON(w, http.StatusOK, user)
}

func (h Handler) Login(w http.ResponseWriter, r *http.Request) error {
	email := r.FormValue("email")
	password := r.FormValue("password")

	// get password by user.email
	user := &model.User{Email: email}
	err := user.GetByEmail(h.Context.Db.PSQLDB)
	if err != nil {
		return h.Context.JSON(w, http.StatusBadRequest, err.Error())
	}

	// compare password and crypt password
	err = user.VerifyPassword(password)
	if err != nil {
		return h.Context.JSON(w, http.StatusUnauthorized, "password is mistaken")
	}

	// create TOKEN
	token := auth.CreateToken(string(user.ID.String()))

	res := LoginResponse{Token: token}
	return h.Context.JSON(w, http.StatusOK, res)
}

func (h Handler) GetUsers(w http.ResponseWriter, r *http.Request) error {

	var u model.User
	users, err := u.GetAll(h.Context.Db.PSQLDB)
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

	err := user.GetByEmail(h.Context.Db.PSQLDB)
	if err != nil {
		return h.Context.JSON(w, http.StatusBadRequest, err.Error())
	}

	return h.Context.JSON(w, http.StatusOK, user)
}

func (h Handler) GetUserBYID(w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)
	id := vars["id"]

	user := &model.User{
		ID: uuid.MustParse(id),
	}

	err := user.GetByUUID(h.Context.Db.PSQLDB)
	if err != nil {
		return h.Context.JSON(w, http.StatusBadRequest, err.Error())
	}

	return h.Context.JSON(w, http.StatusOK, user)
}
