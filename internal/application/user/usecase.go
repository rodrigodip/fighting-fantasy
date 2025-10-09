package userapp

import "github.com/rodrigodip/fighting-fantasy/internal/domain/user"

type UserUseCase struct {
	UserService *user.Service
}

func NewUserUseCase(service *user.Service) *UserUseCase {
	return &UserUseCase{UserService: service}
}

// TODO: REFACTOR: Use case deve ser responsável por orquestrar os serviços
func (uc *UserUseCase) CreateUser(name, email, password string) (*user.User, error) {
	return uc.UserService.CreateUser(name, email, password)
}
func (uc *UserUseCase) GetEmail(email string) (*user.User, error) {
	return uc.UserService.GetEmail(email)
}
