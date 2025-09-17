package authapp

import (
	"errors"
	"fmt"
	"github.com/rodrigodip/fighting-fantasy/internal/domain/user"
	"github.com/rodrigodip/fighting-fantasy/internal/pkg/security"
)

type AuthUseCase struct {
	userRepository user.Repository
	userService    *user.Service
	jwtService     *security.JWTService
}

func NewAuthUseCase(ur user.Repository, us *user.Service, j *security.JWTService) *AuthUseCase {
	return &AuthUseCase{userRepository: ur, userService: us, jwtService: j}
}

func (uc *AuthUseCase) Login(email, password string) (string, error) {
	if err := security.EmailService(email); err != nil {
		return "", fmt.Errorf("E-mail Validation Error: %v", err)
	}

	foundUser, err := uc.userRepository.FindByEmail(email)
	if err != nil || foundUser == nil {
		return "", errors.New("Login: invalid credentials")
	}

	if err := security.CheckPasswordHash(foundUser.Password, password); err != nil {
		return "", errors.New("Login: invalid credentials")
	}

	return uc.jwtService.GenerateToken(foundUser.UserID, foundUser.Roles[0])
}
