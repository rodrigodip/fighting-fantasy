package webmiddleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"github.com/rodrigodip/fighting-fantasy/internal/pkg/security"
)

type WebAuthMiddleware struct {
	sessionStore sessions.Store
	jwtService   *security.JWTService
}

func NewWebAuthMiddleware(store sessions.Store, jwt *security.JWTService) *WebAuthMiddleware {
	return &WebAuthMiddleware{
		sessionStore: store,
		jwtService:   jwt,
	}
}

// RequireAuth checks if user is logged in via session + validates JWT
func (m *WebAuthMiddleware) RequireAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		session, _ := m.sessionStore.Get(c.Request, "session-name")

		token, ok := session.Values["token"].(string)
		if !ok || token == "" {
			c.Redirect(http.StatusSeeOther, "/?error=login_required")
			c.Abort()
			return
		}

		// Validate JWT and extract user info
		userID, err := m.jwtService.ValidateToken(token)
		if err != nil {
			// Token expired or invalid - clear session
			session.Values["token"] = ""
			session.Save(c.Request, c.Writer)
			c.Redirect(http.StatusSeeOther, "/?error=session_expired")
			c.Abort()
			return
		}

		// Set user info in context for handlers
		c.Set("user_id", userID)
		c.Next()
	}
}

// RequireRole checks for specific role
func (m *WebAuthMiddleware) RequireRole(requiredRole string) gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("user_role")
		if !exists {
			c.Redirect(http.StatusSeeOther, "/login")
			c.Abort()
			return
		}

		if role != requiredRole && requiredRole == "admin" {
			// Show forbidden page for web
			c.HTML(http.StatusForbidden, "error.html", gin.H{
				"error": "You don't have permission to access this page",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
