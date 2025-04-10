package handler

import (
	"microservices/api/proto/user"
)

type UserServiceHandler struct {
	user.UserServiceServer
}

func NewUserServiceHandler() *UserServiceHandler {
	return &UserServiceHandler{}
}
