package security

import "golang.org/x/crypto/bcrypt"

// Hash creates a hash from a string
func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

// PwVerify compares a hash password with a string password
func PwVerify(pwHash, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(pwHash), []byte(password))
}
