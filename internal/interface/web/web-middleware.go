package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
)

type WebAuthMiddleware struct {
	sessionStore sessions.Store
}

func NewWebAuthMiddleware(store sessions.Store) *WebAuthMiddleware {
	return &WebAuthMiddleware{
		sessionStore: store,
	}
}

// RequireAuth checks if user is logged in via session
func (m *WebAuthMiddleware) RequireAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		session, _ := m.sessionStore.Get(c.Request, "session-name")

		// Check if user is authenticated
		userID, ok := session.Values["user_id"]
		if !ok {
			// For web: redirect to login
			c.Redirect(http.StatusSeeOther, "/login?redirect="+c.Request.URL.Path)
			c.Abort()
			return
		}

		// Set user info in context for handlers
		c.Set("user_id", userID)
		c.Set("user_role", session.Values["role"])
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
