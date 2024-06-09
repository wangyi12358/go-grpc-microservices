package handler

import (
	"microservices/api/proto/user"
	"microservices/internal/user/model/user_model"
)

type UserServiceHandler struct {
	user.UserServiceServer
	model user_model.UserModel
}

func NewUserServiceHandler() *UserServiceHandler {
	return &UserServiceHandler{
		model: user_model.NewUserModel(),
	}
}

func OfUserRes(u *user_model.User) *user.UserRes {
	return &user.UserRes{
		Id:        u.ID,
		Name:      u.Name,
		Username:  u.Username,
		Password:  u.Password,
		CreatedId: u.CreatedID,
		UpdatedId: u.UpdatedID,
		CreatedAt: u.CreatedAt.String(),
		UpdatedAt: u.UpdatedAt.String(),
	}
}

func OfUsersRes(users []*user_model.User) []*user.UserRes {
	var res []*user.UserRes
	for _, u := range users {
		res = append(res, OfUserRes(u))
	}
	return res
}
