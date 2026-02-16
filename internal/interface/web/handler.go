package webhandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type IWebHandlerRepo interface {
	LoadIndex(c *gin.Context)
}
type WebHandler struct {
	presenter *IWebHandlerRepo
}

func (wh *WebHandler) LoadIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "The FFantasy",
	})
}
