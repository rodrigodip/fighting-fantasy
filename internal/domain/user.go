package domain

import (
	"errors"
	"fmt"
	"github.com/badoux/checkmail"
	"regexp"
)

type User struct {
	ID       string
	Name     string
	Age      int
	Email    string
	Password string
}

func NewUserDomain() *User {
	return &User{}
}

func (u *User) UserValidation() error {
	if u.Name == "" {
		return errors.New("Name Validation Error: requested")
	}
	if len(u.Name) < 3 {
		return errors.New("Name Validation Error: Must have more than 3 digits")
	}

	if u.Email == "" {
		return errors.New("E-mail Validation Error: requested")
	}

	if err := checkmail.ValidateFormat(u.Email); err != nil {
		return fmt.Errorf("E-mail Validation Error: %v", err)
	}

	if u.Password == "" {
		return errors.New("campo senha é obrigatório")
	}
	if err := ValidatePassword(u.Password); err != nil {
		return err
	}
	return nil
}

// ValidatePassword checks if a password is strong
// At least 8 characters long
// Contains at least 1 uppercase letter
// Contains at least 1 lowercase letter
// Contains at least 1 digit
// Contains at least 1 special character (e.g. @#$%^&+=!)
func ValidatePassword(password string) error {
	passwordRegex := regexp.MustCompile(
		`^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[@$!%*?&])[A-Za-z\d@$!%*?&]{8,}$`,
	)
	if passwordRegex.MatchString(password) {
		return nil
	}
	return errors.New("Password Validation Error: Wick Password")
}
