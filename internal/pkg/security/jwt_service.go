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
		"user_id": userID,
		"role":    role,
		"iss":     j.issuer,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(j.secretKey))
}

func (j *JWTService) ValidateToken(token string) (string, error) {
	// secret := os.Getenv("JWT_SECRET")
	// issuer := os.Getenv("JWT_ISSUER")
	t, err := jwt.Parse(token, func(token *jwt.Token) (any, error) {
		return []byte(j.secretKey), nil
	})
	if err != nil || !t.Valid {
	}
	claims := t.Claims.(jwt.MapClaims)
	userID := claims["user_id"].(string)
	return userID, nil
}

// func (j *JWTService) SendVerifyEmail(userID, name, email, role string) error {
// token, err := j.GenerateToken(userID, role)
// if err != nil {
// 	return fmt.Errorf("SendVerifyEmail: GenerateToken: %v", err)
// }
// if err := SendEmail(name, email, token); err != nil {
// 	return fmt.Errorf("SendVerifyEmail: SendEmail: %v", err)
// }
// 	return nil
// }
