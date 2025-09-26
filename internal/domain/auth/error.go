package auth

import (
	"fmt"
)

type AuthErr struct {
	Producer string
	Err      string
}

func (e *AuthErr) Error() string {
	return fmt.Sprintf("%s: %s", e.Producer, e.Err)
}

func InvalidCredentials(producer string) error {
	return &AuthErr{
		Producer: producer,
		Err:      "Invalid Credentials",
	}
}

func NotVerified(producer string) error {
	return &AuthErr{
		Producer: producer,
		Err:      "Email Not Verified",
	}
}
func InvalidToken(producer string) error {
	return &AuthErr{
		Producer: producer,
		Err:      "invalid token",
	}
}
func NotFound(producer string) error {
	return &AuthErr{
		Producer: producer,
		Err:      "user not found",
	}
}
