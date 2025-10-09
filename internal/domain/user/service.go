package user

import (
	"errors"
	"fmt"
	"os"
	"regexp"

	"github.com/rodrigodip/fighting-fantasy/internal/pkg/id_generator"
	"github.com/rodrigodip/fighting-fantasy/internal/pkg/security"
)

type Service struct {
	service Repository
}

func NewUserService(r Repository) *Service {
	return &Service{service: r}
}

// TODO: REFACTOR: Aplicar Responsabilidades únicas aos serviços.
func (s Service) GetEmail(email string) (*User, error) {
	if err := security.EmailService(email); err != nil {
		return &User{}, fmt.Errorf("E-mail Validation Error: %v", err)
	}
	foudUser, err := s.service.FindByEmail(email)
	if err != nil {
		return &User{}, err
	}

	return foudUser, nil
}
func (s *Service) CreateUser(name, email, password string) (*User, error) {
	if name == "" {
		return &User{}, errors.New("Name Validation Error: Requested")
	}
	if len(name) < 3 {
		return &User{}, errors.New("Name Validation Error: Must have more than 3 digits")
	}

	if email == "" {
		return &User{}, errors.New("E-mail Validation Error: Requested")
	}

	if err := security.EmailService(email); err != nil {
		return &User{}, fmt.Errorf("E-mail Validation Error: %v", err)
	}

	if password == "" {
		return &User{}, errors.New("Password Validation Error: Requested")
	}
	if err := validatePassword(password); err != nil {
		return &User{}, err
	}
	encodedPassword, err := security.EncodePw(password)
	if err != nil {
		return &User{}, err
	}
	generator := IDgenerator.NewTimestampIDGenerator()
	newUser := &User{
		UserID:   generator.NewID(),
		Name:     name,
		Email:    email,
		Password: encodedPassword,
		Role:     "USER",
		Status:   "UNVERIFIED",
	}
	if err := s.service.RegisterUser(
		newUser.UserID,
		newUser.Name,
		newUser.Email,
		newUser.Password,
		newUser.Role,
		newUser.Status); err != nil {
		return &User{}, err
	}
	sendVerifyEmail(newUser.UserID, newUser.Name, newUser.Email)
	return newUser, nil
}

//NOTE:
// At least 8 characters long
// Contains at least 1 uppercase letter
// Contains at least 1 lowercase letter
// Contains at least 1 digit
// Contains at least 1 special character (e.g. @#$%^&+=!)

// ValidatePassword checks if a password is strong
func validatePassword(password string) error {
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
func sendVerifyEmail(userID, name, email string) error {
	role := "USER"
	secret := os.Getenv("JWT_SECRET")
	issuer := os.Getenv("JWT_ISSUER")
	service := security.NewJWTService(secret, issuer)
	token, err := service.GenerateToken(userID, role)
	if err != nil {
		return errors.New("Error generating email validation Token")
	}
	security.SendEmail(name, email, token)
	return nil
}
