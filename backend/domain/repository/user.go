package repository

import "github.com/yoshihiro-shu/draft-backend/backend/domain/model"

type UserRepository interface {
	Create(user *model.User) error
	FindByID(id int) (*model.User, error)
	FindByEmail(email string) (*model.User, error)
	Update(user *model.User) error
	Delete(user *model.User) error
}
