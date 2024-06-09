package user_model

import (
	"gorm.io/gorm"
	"microservices/pkg/model"
)

type UserModel interface {
	UserList() ([]*User, error)
	Total() (int64, error)
}

type userModel struct {
	db *gorm.DB
}

func NewUserModel() UserModel {
	return &userModel{
		db: model.DB,
	}
}

func (u *userModel) UserList() ([]*User, error) {
	var users []*User
	err := model.DB.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (u *userModel) Total() (int64, error) {
	var total int64
	err := model.DB.Model(&User{}).Count(&total).Error
	if err != nil {
		return 0, err
	}
	return total, nil
}
