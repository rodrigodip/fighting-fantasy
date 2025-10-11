package user

import (
	"fmt"
)

type UserErr struct {
	Producer string
	Err      string
}

func (e *UserErr) Error() string {
	return fmt.Sprintf("%s: %s", e.Producer, e.Err)
}

// InvalidCredentials generates a UserErr type and indicates where the error producer
func InvalidCredentials(producer string) error {
	return &UserErr{
		Producer: producer,
		Err:      "Invalid Credentials",
	}
}
func NotVerified(producer string) error {
	return &UserErr{
		Producer: producer,
		Err:      "Email Not Verified",
	}
}
func InvalidToken(producer string) error {
	return &UserErr{
		Producer: producer,
		Err:      "invalid token",
	}
}
func NotFound(producer string) error {
	return &UserErr{
		Producer: producer,
		Err:      "user not found",
	}
}
