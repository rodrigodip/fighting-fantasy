package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rodrigodip/fighting-fantasy/internal/infrastructure/di"
)

func InitWebHeroGroup(r *gin.RouterGroup, d *dependecy.Container) {
	// Protected routes - require authentication
	protected := r.Group("/")
	protected.Use(d.WebAuthMiddleware.RequireAuth())
	{
		protected.GET("/dashboard", nil)    //d.WebHeroHandlers.Dashboard)
		protected.GET("/hero/create", nil)  //d.WebHeroHandlers.ShowCreateForm)
		protected.POST("/hero/create", nil) //d.WebHeroHandlers.CreateHero)
	}
}
func InitWebAuthGroup(r *gin.RouterGroup, d *dependecy.Container) {
	// Public routes - no auth required
	r.GET("/", d.WebUserHandlers.LandingPage)
	r.GET("/login-form", d.WebUserHandlers.ShowLoginForm)
	r.POST("/login", d.WebUserHandlers.Login)
	r.GET("/register-form", d.WebUserHandlers.ShowRegisterForm)
	r.POST("/register", d.WebUserHandlers.CreateUserFromWeb)
	r.GET("/logout", d.WebUserHandlers.Logout)
}
