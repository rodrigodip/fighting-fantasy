package auth

import (
	"api/internal/config"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// NewToken generates permission token for a user by id
func NewToken(userId uint64) (string, error) {
	permisions := jwt.MapClaims{}
	permisions["autorized"] = true
	permisions["exp"] = time.Now().Add(time.Hour * 6).Unix()
	permisions["userId"] = userId
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permisions)
	return token.SignedString([]byte(config.SecreteKey)) //secret
}
