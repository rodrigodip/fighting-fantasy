package auth

import (
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
		return "", InvalidCredentials("SRV-0001")
	}

	foundUser, err := uc.userRepository.FindByEmail(email)
	if err != nil || foundUser == nil {
		return "", InvalidCredentials("SRV-0010")
	}

	if foundUser.Status != "VERIFIED" {
		return "", NotVerified("SRV-0100")
	}
	if err := security.CheckPasswordHash(foundUser.Password, password); err != nil {
		return "", InvalidCredentials("SRV-1000")
	}
	return uc.jwtService.GenerateToken(foundUser.UserID, foundUser.Role)
}
func (uc *Service) VerifyEmail(token string) error {
	t, err := uc.jwtService.ValidateToken(token)
	if err != nil || !t.Valid {
		return InvalidToken("SRV-1010")
	}

	claims := t.Claims.(jwt.MapClaims)
	userID := claims["user_id"].(string)

	foundUser, err := uc.userRepository.FindById(userID)
	if err != nil || foundUser == nil {
		return NotFound("SRV-1001")
	}
	if err = uc.userRepository.SetVerified(userID); err != nil {
		return err
	}

	return nil
}
