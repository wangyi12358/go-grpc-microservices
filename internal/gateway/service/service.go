package service

import "microservices/api/proto/user"

type Service struct {
	UserService user.UserServiceClient
}

func New(user user.UserServiceClient) *Service {
	return &Service{
		UserService: user,
	}
}
