package webhandlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *HeroHandler) HomePage(c *gin.Context) {
	log.Print("Home Page")
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "The forest of Destruction",
	})
}
