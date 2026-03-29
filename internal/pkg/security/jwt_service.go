package security

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTService struct {
	secretKey string
	issuer    string
}

// TokenClaims holds the parsed data your middleware and handlers need
type TokenClaims struct {
	UserID string
	Role   string
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

// ValidateToken kept as-is — preserves compatibility with any existing callers
func (j *JWTService) ValidateToken(token string) (string, error) {
	claims, err := j.parseAndValidate(token)
	if err != nil {
		return "", err
	}
	return claims.UserID, nil
}

// ValidateTokenWithClaims returns both userID and role — use this in middleware
func (j *JWTService) ValidateTokenWithClaims(token string) (*TokenClaims, error) {
	return j.parseAndValidate(token)
}

// parseAndValidate is the single place where JWT parsing lives
func (j *JWTService) parseAndValidate(tokenStr string) (*TokenClaims, error) {
	t, err := jwt.Parse(tokenStr, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(j.secretKey), nil
	})
	if err != nil || !t.Valid {
		return nil, fmt.Errorf("invalid token: %w", err)
	}

	mapClaims, ok := t.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("invalid claims format")
	}

	userID, ok := mapClaims["userId"].(string)
	if !ok {
		return nil, fmt.Errorf("userId claim missing or not a string")
	}

	role, ok := mapClaims["role"].(string)
	if !ok {
		return nil, fmt.Errorf("role claim missing or not a string")
	}

	return &TokenClaims{
		UserID: userID,
		Role:   role,
	}, nil
}

