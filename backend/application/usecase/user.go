package usecase

import (
	"github.com/yoshihiro-shu/draft-backend/domain/model"
	"github.com/yoshihiro-shu/draft-backend/domain/repository"
	"github.com/yoshihiro-shu/draft-backend/interfaces/api/server/auth"
)

type UserUseCase interface {
	Create(name, password, email string) error
	Login(email, password string) (*auth.AuthToken, error)
	Update(id int, name, password, email string) error
	Delete(id int) error
}

type userUseCase struct {
	userRepo         repository.UserRepository
	refreshTokenRepo repository.RefreshTokenRepository
}

func NewUserUseCase(userRepo repository.UserRepository) UserUseCase {
	return &userUseCase{userRepo: userRepo}
}

func (uu *userUseCase) Create(name, password, email string) error {
	user := model.NewUser(name, password, email)

	err := uu.userRepo.Create(user)
	if err != nil {
		return err
	}

	return nil
}

func (uu *userUseCase) FindByID(id int) (*model.User, error) {
	user, err := uu.userRepo.FindByID(id)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (uu *userUseCase) FindByEmail(email string) (*model.User, error) {
	user, err := uu.userRepo.FindByEmail(email)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *userUseCase) Login(email, password string) (*auth.AuthToken, error) {
	user, err := u.userRepo.FindByEmail(email)
	if err != nil {
		return nil, err
	}

	if !auth.IsVerifyPassword(password, user.Password) {
		return nil, auth.ErrInvalidPassword
	}

	accessToken := auth.CreateAccessToken(user.Id)

	refreshToken := auth.GenerateToken()
	err = u.refreshTokenRepo.Create(user.Id, refreshToken)
	if err != nil {
		return nil, err
	}

	return &auth.AuthToken{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (uu *userUseCase) Update(id int, name, password, email string) error {
	user := &model.User{
		Id:       id,
		Name:     name,
		Password: password,
		Email:    email,
	}

	err := uu.userRepo.Update(user)
	if err != nil {
		return err
	}

	return err
}

func (uu *userUseCase) Delete(id int) error {
	user := &model.User{
		Id: id,
	}
	err := uu.userRepo.Delete(user)
	if err != nil {
		return err
	}

	return nil
}
