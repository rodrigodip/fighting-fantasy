package security

import "golang.org/x/crypto/bcrypt"

// Encode creates a hash from a string
func EncodePw(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

// Verify compares a hash password with a string password
func CheckPasswordHash(pwHash []byte, password string) error {
	return bcrypt.CompareHashAndPassword(pwHash, []byte(password))
}
