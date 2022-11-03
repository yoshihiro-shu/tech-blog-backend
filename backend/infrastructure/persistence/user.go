package persistence

import (
	"github.com/go-pg/pg"
	"github.com/yoshihiro-shu/draft-backend/domain/model"
	"github.com/yoshihiro-shu/draft-backend/domain/repository"
)

type userPersistence struct {
	Conn *pg.DB
}

func NewUserPersistence(conn *pg.DB) repository.UserRepository {
	return &userPersistence{Conn: conn}
}

func (up *userPersistence) Create(user *model.User) error {
	_, err := up.Conn.Model(user).Insert()
	if err != nil {
		return err
	}

	return nil
}

func (up *userPersistence) FindByID(id int) (*model.User, error) {
	user := &model.User{Id: id}
	query := up.Conn.Model(user).WherePK()

	err := query.Select()
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (up *userPersistence) FindByEmail(email string) (*model.User, error) {
	user := &model.User{}

	query := up.Conn.Model(user).Where("email = ?", email)
	err := query.Select()
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (up *userPersistence) Update(user *model.User) error {
	_, err := up.Conn.Model(user).WherePK().Update()
	if err != nil {
		return err
	}
	return nil
}

func (up *userPersistence) Delete(user *model.User) error {
	_, err := up.Conn.Model(user).Delete()
	if err != nil {
		return nil
	}
	return nil
}
