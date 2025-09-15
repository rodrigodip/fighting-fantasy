package userapp

import "github.com/rodrigodip/fighting-fantasy/internal/domain/user"

type UserUseCase struct {
	UserService *user.Service
}

func NewUserUseCase(service *user.Service) *UserUseCase {
	return &UserUseCase{UserService: service}
}

func (uc *UserUseCase) CreateUser(name, email, password string) (*user.User, error) {
	return uc.UserService.CreateUser(name, email, password)
}
func (uc *UserUseCase) GetEmail(email string) (*user.User, error) {
	return uc.UserService.GetEmail(email)
}
