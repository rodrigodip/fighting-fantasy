package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rodrigodip/fighting-fantasy/internal/infrastructure/di"
)

func InitWebPagesGroup(r *gin.RouterGroup, d *dependecy.Container) {
	// Public pages — no auth required
	r.GET("/", d.WebPagesHandlers.AuthPageHandler)
	r.GET("/verify", d.WebPagesHandlers.EmailVerifyPageHandler)

	// Protected pages — session required
	protected := r.Group("/")
	protected.Use(d.WebAuthMiddleware.RequireAuth())
	{
		protected.GET("/dashboard", d.WebPagesHandlers.DashboardPageHandler)
		protected.GET("/adventure", d.WebPagesHandlers.AdventurePageHandler)
		// protected.GET("/gameover",  d.WebPagesHandlers.GameOverPageHandler)
	}
}

func InitWebAuthGroup(r *gin.RouterGroup, d *dependecy.Container) {
	// All public — no auth required
	r.POST("/auth/login", d.WebUserHandlers.AuthLoginHandler)
	r.POST("/auth/signup", d.WebUserHandlers.AuthSignUpHandler)
	r.POST("/auth/logout", d.WebUserHandlers.AuthLogoutHandle)
	r.GET("/auth/verify", d.WebUserHandlers.AuthVerifyEmailHandler)
}

func InitWebHeroGroup(r *gin.RouterGroup, d *dependecy.Container) {
	protected := r.Group("/hero")
	protected.Use(d.WebAuthMiddleware.RequireAuth())
	{
		protected.POST("/create", d.WebHeroHandlers.HeroCreateHandler)
	}
}
