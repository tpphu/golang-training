package repo

import (
	"../model"
	"github.com/jinzhu/gorm"
)

type UserRepo interface {
	Create(model.User) (*model.User, error)
	FindByUserLogin(string) (*model.User, error)
}

type UserRepoImpl struct {
	DB *gorm.DB
}

func (self *UserRepoImpl) Create(user model.User) (*model.User, error) {
	err := self.DB.Create(&user).Error
	return &user, err
}

func (self *UserRepoImpl) FindByUserLogin(login string) (*model.User, error) {
	user := &model.User{}
	err := self.DB.
		Where("username = ? OR email = ?", login, login).
		First(user).Error
	return user, err
}
