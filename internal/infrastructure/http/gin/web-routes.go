package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rodrigodip/fighting-fantasy/internal/infrastructure/di"
)

func InitViewGroup(r *gin.RouterGroup, d *dependecy.Container) {
	r.GET("/testhero", d.WebHero.Detail)
	r.GET("/", d.WebHero.HomePage)
}
