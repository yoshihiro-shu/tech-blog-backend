package persistence

import (
	"github.com/yoshihiro-shu/draft-backend/backend/domain/model"
	"github.com/yoshihiro-shu/draft-backend/backend/domain/repository"
	"gorm.io/gorm"
)

type userPersistence struct {
	Master  func() *gorm.DB
	Reprica func() *gorm.DB
}

func NewUserPersistence(master, reprica func() *gorm.DB) repository.UserRepository {
	return &userPersistence{
		Master:  master,
		Reprica: reprica,
	}
}

func (up *userPersistence) Create(user *model.User) error {
	return up.Master().Model(&model.User{}).Create(user).Error
}

func (up *userPersistence) FindByID(id int) (*model.User, error) {
	user := &model.User{Id: id}
	err := up.Reprica().Model(&model.User{}).Find(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (up *userPersistence) FindByEmail(email string) (*model.User, error) {
	user := &model.User{}
	err := up.Reprica().Model(&model.User{}).Where("email = ?", email).Find(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (up *userPersistence) Update(user *model.User) error {
	return up.Master().Model(&model.User{}).Where("id = ?", user.Id).Updates(user).Error
}

func (up *userPersistence) Delete(user *model.User) error {
	return up.Master().Model(&model.User{}).Where("id = ?", user.Id).Delete(user).Error
}
