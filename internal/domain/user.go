package domain

import (
	"errors"
	"fmt"
	"github.com/badoux/checkmail"
	"regexp"
)

type Profile struct {
	Age    int
	Gender string
}

type User struct {
	UserID   string
	Name     string
	Email    string
	Password string
	Status   string
	Profile  Profile
	Roles    []string
}

func NewUserDomain() *User {
	return &User{}
}

func (u *User) UserValidation() error {
	if u.Name == "" {
		return errors.New("Name Validation Error: Requested")
	}
	if len(u.Name) < 3 {
		return errors.New("Name Validation Error: Must have more than 3 digits")
	}

	if u.Email == "" {
		return errors.New("E-mail Validation Error: Requested")
	}

	if err := checkmail.ValidateFormat(u.Email); err != nil {
		return fmt.Errorf("E-mail Validation Error: %v", err)
	}

	if u.Password == "" {
		return errors.New("Password Validation Error: Requested")
	}
	if err := ValidatePassword(u.Password); err != nil {
		return err
	}
	if len(u.Roles) < 1 {
		return errors.New("Roles Validation Error: Requested")
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
	//return errors.New("Password Validation Error: Wick Password")
}
