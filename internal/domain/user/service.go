package user

import (
	"errors"
	"fmt"
	"regexp"

	"github.com/badoux/checkmail"
	"github.com/rodrigodip/fighting-fantasy/internal/pkg/id_generator"
	"github.com/rodrigodip/fighting-fantasy/internal/pkg/password"
)

type Service struct {
	service Repository
}

func NewUserService(r Repository) *Service {
	return &Service{service: r}
}
func (s Service) GetEmail(email string) (*User, error) {
	if err := checkmail.ValidateFormat(email); err != nil {
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

	if err := checkmail.ValidateFormat(email); err != nil {
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
	newUser := &User{
		UserID:   IDgenerator.NewSimpleID(),
		Name:     name,
		Email:    email,
		Password: encodedPassword,
		Roles:    []string{"user"},
	}
	if err := s.service.RegisterUser(
		newUser.UserID,
		newUser.Name,
		newUser.Email,
		newUser.Password,
		newUser.Roles); err != nil {
		return &User{}, err
	}
	return newUser, nil
}

// ValidatePassword checks if a password is strong
// At least 8 characters long
// Contains at least 1 uppercase letter
// Contains at least 1 lowercase letter
// Contains at least 1 digit
// Contains at least 1 special character (e.g. @#$%^&+=!)
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
