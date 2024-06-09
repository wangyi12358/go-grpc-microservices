package handler

import (
	"context"
	"microservices/api/proto/user"
)

func (u *UserServiceHandler) UserPage(context.Context, *user.UserPageReq) (*user.UserPageRes, error) {
	return &user.UserPageRes{
		Total: 10,
	}, nil
}
