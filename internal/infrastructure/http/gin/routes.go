package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rodrigodip/fighting-fantasy/internal/domain/auth"
	"github.com/rodrigodip/fighting-fantasy/internal/infrastructure/config"
	"github.com/rodrigodip/fighting-fantasy/internal/interface/http_handler"
	"github.com/rodrigodip/fighting-fantasy/internal/interface/http_handler/auth"
	"github.com/rodrigodip/fighting-fantasy/internal/interface/http_handler/hero"
	"github.com/rodrigodip/fighting-fantasy/internal/interface/http_handler/user"
)

func InitUserGroup(r *gin.RouterGroup, app userhandler.UserHandlerRepo) {
	r.POST("/users", app.RegisterUser)

	//TODO: Swagger Doc
	//docs.SwaggerInfo.BasePath = "/"
	//r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
func InitAuthGroup(r *gin.RouterGroup, app authhandler.AuthHandlerRepo) {
	r.POST("/login", app.Login)
	r.GET("/verify", app.VerifyEmail)
}
func InitHeroGroup(r *gin.RouterGroup, app herohandler.HeroHandlerRepo, cfg config.Config) {
	r.POST("/heroes", interfaces.AuthMiddleware(cfg.JWTSecret, auth.RoleUser), app.RegisterHero)
}
