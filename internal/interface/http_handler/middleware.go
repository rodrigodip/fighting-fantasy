package interfaces

import (
	"net/http"
	"strings"

	"github.com/rodrigodip/fighting-fantasy/internal/domain/user"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware(secret string, requiredRole usr.Role) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "missing token"})
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (any, error) {
			return []byte(secret), nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			c.Abort()
			return
		}

		claims := token.Claims.(jwt.MapClaims)
		role := claims["role"].(string)

		if requiredRole == usr.RoleAdmin && role != string(usr.RoleAdmin) { //!usr.CanAccessAdmin(usr.Role(role))
			c.JSON(http.StatusForbidden, gin.H{"error": "Must have admin access"})
			c.Abort()
			return
		}

		c.Set("userId", claims["userId"])
		c.Set("role", role)
		c.Next()
	}
}
