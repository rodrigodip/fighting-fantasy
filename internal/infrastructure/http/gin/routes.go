package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rodrigodip/fighting-fantasy/internal/domain/user"
	"github.com/rodrigodip/fighting-fantasy/internal/infrastructure/config"
	"github.com/rodrigodip/fighting-fantasy/internal/interface/http_handler"
	"github.com/rodrigodip/fighting-fantasy/internal/interface/http_handler/hero"
	"github.com/rodrigodip/fighting-fantasy/internal/interface/http_handler/user"
	docs "github.com/rodrigodip/fighting-fantasy/internal/pkg/swagger/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitUserGroup(r *gin.RouterGroup, app *userhandler.UserHandler, cfg config.Config) {
	r.POST("/users", app.RegisterUser)
	r.POST("/login", app.Login)
	r.GET("/verify", app.VerifyEmail)

	r.POST("/heroes", interfaces.AuthMiddleware(cfg.JWTSecret, usr.RoleUser), app.RegisterHero)

	docs.SwaggerInfo.BasePath = "/"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
func InitHeroGroup(r *gin.RouterGroup, app *herohandler.HeroHandler, cfg config.Config) {
	// r.POST("/heroes", interfaces.AuthMiddleware(cfg.JWTSecret, usr.RoleUser), app.RegisterHero)
}
