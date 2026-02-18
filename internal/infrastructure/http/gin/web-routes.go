package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rodrigodip/fighting-fantasy/internal/infrastructure/di"
)

func InitWebHeroGroup(r *gin.RouterGroup, d *dependecy.Container) {
	r.GET("/testhero", nil)
}

func InitWebUserGroup(r *gin.RouterGroup, d *dependecy.Container) {
	r.GET("/register", d.WebUserHandlers.ShowRegisterForm)
	r.POST("/register", d.WebUserHandlers.CreateUserFromWeb) // Handle form submission
}
