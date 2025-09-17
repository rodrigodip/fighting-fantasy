package security

import "github.com/badoux/checkmail"

func EmailService(email string) error {
	err := checkmail.ValidateFormat(email)
	return err
}
