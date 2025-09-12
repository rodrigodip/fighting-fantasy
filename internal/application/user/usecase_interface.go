package userapp

import "github.com/rodrigodip/fighting-fantasy/internal/domain/user"

type UserUseCase struct {
	UserService *user.Service
}

func NewUserUseCase(service *user.Service) *UserUseCase {
	return &UserUseCase{UserService: service}
}
