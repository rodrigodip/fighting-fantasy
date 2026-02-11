package usr

import (
	"errors"
	"fmt"
	"regexp"

	"github.com/rodrigodip/fighting-fantasy/internal/pkg/security"
)

type Service struct{}

func NewUserService() *Service {
	return &Service{}
}

// ValidadeUserInput enforce business rules
func (s *Service) ValidadeUserInput(name, email, password string) error {
	if name == "" {
		return errors.New("Name Validation Error: Requested")
	}
	if len(name) < 3 {
		return errors.New("Name Validation Error: Must have more than 3 digits")
	}
	if email == "" {
		return errors.New("E-mail Validation Error: Requested")
	}
	if err := security.EmailFormatValidation(email); err != nil {
		return fmt.Errorf("E-mail Validation Error: %v", err)
	}
	if password == "" {
		return errors.New("Password Validation Error: Requested")
	}
	return nil
}

//NOTE:
// At least 8 characters long
// Contains at least 1 uppercase letter
// Contains at least 1 lowercase letter
// Contains at least 1 digit
// Contains at least 1 special character (e.g. @#$%^&+=!)

// ValidatePassword checks if a password rules
func (s *Service) ValidatePassword(password string) error {
	if len(password) < 8 {
		return errors.New("Password Validation Error: Must be at least 8 characters long")
	}
	if !regexp.MustCompile(`[a-z]`).MatchString(password) {
		return errors.New("Password Validation Error: Must contain at least one lowercase letter")
	}

	if !regexp.MustCompile(`[A-Z]`).MatchString(password) {
		return errors.New("Password Validation Error: Must contain at least one uppercase letter")
	}

	if !regexp.MustCompile(`\d`).MatchString(password) {
		return errors.New("Password Validation Error: Must contain at least one digit")
	}

	if !regexp.MustCompile(`[@$!%*?&]`).MatchString(password) {
		return errors.New("Password Validation Error: Must contain at least one special character (@$!%*?&)")
	}
	return nil
}

// func (s *Service) GetToken(userID, role string) (string, error) {
// 	secret := os.Getenv("JWT_SECRET")
// 	issuer := os.Getenv("JWT_ISSUER")
// 	repository := security.NewJWTService(secret, issuer)
// 	token, err := service.GenerateToken(userID, role)
// 	if err != nil {
// 		return "", authErr.InvalidCredentials("SRV-0001")
// 	}
// 	return token, nil
// }
// func (s *Service) CheckToken(token string) (*jwt.Token, error) {
// 	secret := os.Getenv("JWT_SECRET")
// 	issuer := os.Getenv("JWT_ISSUER")
// 	JWTrepository := security.NewJWTService(secret, issuer)
// 	t, err := JWTservice.ValidateToken(token)
// 	if err != nil || !t.Valid {
// 		return &jwt.Token{}, authErr.InvalidToken("SRV-1010")
// 	}
// 	return t, nil
// }

// SendVerifyEmail sends a verification link to user registered email
// func (s *Service) SendVerifyEmail(userID, name, email, role string) error {
// 	token, err := s.auth.GenerateToken(userID, role)
// 	if err != nil {
// 		return fmt.Errorf("SendVerifyEmail: GenerateToken: %v", err)
// 	}
// 	security.SendEmail(name, email, token)
// 	return nil
// }

// IsUserVerified check if User has a verified e-mail
func (s *Service) IsUserVerified(status Status) bool {
	return status == StatusUnverified
}
