package webmiddleware

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"github.com/rodrigodip/fighting-fantasy/internal/pkg/security"
)

const sessionName = "user-session" // single source of truth — use everywhere

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

// RequireAuth checks if user is logged in via session + validates JWT.
// Sets "user_id" and "user_role" in context for downstream handlers.
func (m *WebAuthMiddleware) RequireAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		session, err := m.sessionStore.Get(c.Request, sessionName)
		if err != nil {
			log.Printf("DEBUG session get err: %v", err)
			c.Redirect(http.StatusSeeOther, "/?error=session_error")
			c.Abort()
			return
		}

		token, ok := session.Values["token"].(string)
		log.Printf("DEBUG token ok: %v | token empty: %v", ok, token == "")

		if !ok || token == "" {
			log.Printf("DEBUG no token in session — redirecting")
			c.Redirect(http.StatusSeeOther, "/?error=login_required")
			c.Abort()
			return
		}

		claims, err := m.jwtService.ValidateTokenWithClaims(token)
		log.Printf("DEBUG validate err: %v | claims: %+v", err, claims)

		if err != nil {
			session.Values["token"] = ""
			session.Save(c.Request, c.Writer)
			c.Redirect(http.StatusSeeOther, "/?error=session_expired")
			c.Abort()
			return
		}

		c.Set("user_id", claims.UserID)
		c.Set("user_role", claims.Role)
		c.Next()
	}
}

// RequireRole must be chained AFTER RequireAuth in the middleware stack.
func (m *WebAuthMiddleware) RequireRole(requiredRole string) gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("user_role")
		if !exists {
			// RequireAuth wasn't in the chain — programming error, not user error
			c.HTML(http.StatusInternalServerError, "error.html", gin.H{
				"error": "Auth middleware misconfiguration",
			})
			c.Abort()
			return
		}
		if role != requiredRole {
			c.HTML(http.StatusForbidden, "error.html", gin.H{
				"error": "You don't have permission to access this page",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
