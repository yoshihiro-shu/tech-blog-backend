package usecase

import (
	"github.com/yoshihiro-shu/tech-blog-backend/backend/domain/model"
	"github.com/yoshihiro-shu/tech-blog-backend/backend/domain/repository"
	"github.com/yoshihiro-shu/tech-blog-backend/backend/internal/auth"
)

type UserUseCase interface {
	Create(name, password, email string) error
	Login(email, password string) (*auth.AuthToken, error)
	RefreshToken(refreshToken string) (*auth.AuthToken, error)
	Update(id int, name, password, email string) error
	Delete(id int) error
}

type userUseCase struct {
	userRepo         repository.UserRepository
	refreshTokenRepo repository.RefreshTokenRepository
}

func NewUserUseCase(userRepo repository.UserRepository, refreshTokenRepo repository.RefreshTokenRepository) UserUseCase {
	return &userUseCase{
		userRepo:         userRepo,
		refreshTokenRepo: refreshTokenRepo,
	}
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

func (u *userUseCase) Login(email, password string) (*auth.AuthToken, error) {
	user, err := u.userRepo.FindByEmail(email)
	if err != nil {
		return nil, err
	}

	if !auth.IsVerifyPassword(password, user.Password) {
		return nil, auth.ErrInvalidPassword
	}

	accessToken := auth.NewAccessToken(user.Id)
	refreshToken := auth.NewRefreshToken(user.Id)

	err = u.refreshTokenRepo.Create(refreshToken.UserId, refreshToken.JwtId, refreshToken.ExpiredAt)
	if err != nil {
		return nil, err
	}

	return &auth.AuthToken{
		AccessToken:  *accessToken,
		RefreshToken: *refreshToken,
	}, nil
}

func (u *userUseCase) RefreshToken(token string) (*auth.AuthToken, error) {
	refreshToken, err := auth.VerifyRefeshToken(token)
	if err != nil {
		return nil, err
	}

	rt, err := u.refreshTokenRepo.GetByJwtId(refreshToken.JwtId)
	if err != nil {
		return nil, err
	}

	newAccessToken := auth.NewAccessToken(rt.UserId)
	newRefreshToken := auth.NewRefreshToken(rt.UserId)
	err = u.refreshTokenRepo.Update(rt.Id, newRefreshToken.JwtId, newRefreshToken.ExpiredAt)
	if err != nil {
		return nil, err
	}

	return &auth.AuthToken{
		AccessToken:  *newAccessToken,
		RefreshToken: *newRefreshToken,
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
