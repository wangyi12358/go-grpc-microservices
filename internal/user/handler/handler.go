package handler

import (
	"microservices/api/proto/user"
)

type UserServiceHandler struct {
	user.UserServiceServer
}
