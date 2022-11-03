package usecase

import (
	"github.com/yoshihiro-shu/draft-backend/domain/model"
	"github.com/yoshihiro-shu/draft-backend/domain/repository"
)

type UserUseCase interface {
	Create(name, password, email string) error
	FindByID(id int) (*model.User, error)
	FindByEmail(email string) (*model.User, error)
	Update(id int, name, password, email string) error
	Delete(id int) error
}

type userUseCase struct {
	userRepo repository.UserRepository
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
