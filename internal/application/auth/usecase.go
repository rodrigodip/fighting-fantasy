package authapp

import (
	"github.com/rodrigodip/fighting-fantasy/internal/domain/auth"
)

type AuthUseCase struct {
	AuthService *auth.Service
}

// TODO: REFACTOR: Atribuir os UC de Auth ao User
func NewAuthUseCase(service *auth.Service) *AuthUseCase {
	return &AuthUseCase{AuthService: service}
}
func (uc *AuthUseCase) VerifyEmail(token string) error {
	return uc.AuthService.VerifyEmail(token)
}
func (uc *AuthUseCase) Login(email, password string) (string, error) {
	return uc.AuthService.Login(email, password)
}
