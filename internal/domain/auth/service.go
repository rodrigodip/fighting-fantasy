package auth

import (
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
	"github.com/rodrigodip/fighting-fantasy/internal/domain/user"
	"github.com/rodrigodip/fighting-fantasy/internal/pkg/security"
)

type Service struct {
	userRepository user.Repository
	userService    *user.Service
	jwtService     *security.JWTService
}

func NewAuthService(ur user.Repository, us *user.Service, j *security.JWTService) *Service {
	return &Service{userRepository: ur, userService: us, jwtService: j}
}

func (uc *Service) Login(email, password string) (string, error) {
	if err := security.EmailService(email); err != nil {
		return "", fmt.Errorf("E-mail Validation Error: %v", err)
	}

	foundUser, err := uc.userRepository.FindByEmail(email)
	if err != nil || foundUser == nil {
		return "", errors.New("Login: invalid credentials")
	}

	if foundUser.Status != "VERIFIED" {
		return "", errors.New("Login: email not verified")
	}

	if err := security.CheckPasswordHash(foundUser.Password, password); err != nil {
		return "", errors.New("Login: invalid credentials")
	}

	return uc.jwtService.GenerateToken(foundUser.UserID, foundUser.Role)
}
func (uc *Service) VerifyEmail(token string) error {
	t, err := uc.jwtService.ValidateToken(token)
	if err != nil || !t.Valid {
		return errors.New("invalid token")
	}

	claims := t.Claims.(jwt.MapClaims)
	userID := claims["user_id"].(string)

	foundUser, err := uc.userRepository.FindById(userID)
	if err != nil || foundUser == nil {
		return errors.New("user not found")
	}
	if err = uc.userRepository.SetVerified(userID); err != nil {
		return err
	}

	return nil
}

