package handler

import (
	"context"
	"microservices/api/proto/user"
)

func (u *UserServiceHandler) UserPage(context.Context, *user.UserPageReq) (*user.UserPageRes, error) {
	users, err := u.model.UserList()
	if err != nil {
		return nil, err
	}
	total, err := u.model.Total()
	if err != nil {
		return nil, err
	}
	return &user.UserPageRes{
		List:  OfUsersRes(users),
		Total: int32(total),
	}, nil
}
