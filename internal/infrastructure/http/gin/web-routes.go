package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rodrigodip/fighting-fantasy/internal/infrastructure/di"
)

// Page routes - render HTML templates
func InitWebPagesGroup(r *gin.RouterGroup, d *dependecy.Container) {
	// TODO: Uncomment and implement when handlers are ready
	// http.HandleFunc("/", handlers.AuthPageHandler)
	// http.HandleFunc("/dashboard", handlers.DashboardPageHandler)
	// http.HandleFunc("/adventure", handlers.AdventurePageHandlerView)
	// http.HandleFunc("/gameover", handlers.GameOverPageHandler)
}

func InitWebAuthGroup(r *gin.RouterGroup, d *dependecy.Container) {
	// Public routes - no auth required
	r.POST("/auth/login", d.WebUserHandlers.AuthLoginHandler)
	r.POST("/auth/signup", d.WebUserHandlers.AuthSignUpHandler)
	r.POST("/auth/logout", d.WebUserHandlers.AuthLogoutHandle)
	r.GET("/auth/verify", d.WebUserHandlers.AuthVerifyEmailHandler)
}
