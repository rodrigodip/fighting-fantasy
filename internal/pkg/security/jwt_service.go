package security

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTService struct {
	secretKey string
	issuer    string
}

func NewJWTService(secret, issuer string) *JWTService {
	return &JWTService{secretKey: secret, issuer: issuer}
}

func (j *JWTService) GenerateToken(userID, role string) (string, error) {
	claims := jwt.MapClaims{
		"userId": userID,
		"role":   role,
		"iss":    j.issuer,
		"exp":    time.Now().Add(24 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(j.secretKey))
}

func (j *JWTService) ValidateToken(token string) (string, error) {
	t, err := jwt.Parse(token, func(token *jwt.Token) (any, error) {
		return []byte(j.secretKey), nil
	})
	if err != nil || !t.Valid {
		return "", err
	}
	claims := t.Claims.(jwt.MapClaims)
	userID := claims["userId"].(string)
	return userID, nil
}
