package herohandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rodrigodip/fighting-fantasy/internal/application/hero"
)

type HeroHandlerRepo interface {
	RegisterHero(c *gin.Context)
}

type HeroHandler struct {
	usecase *heroapp.HeroUseCase
}

func NewHeroHandler(uc *heroapp.HeroUseCase) *HeroHandler {
	return &HeroHandler{usecase: uc}
}

func (hh *HeroHandler) RegisterHero(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "missing token"})
		c.Abort()
		return
	}
	var req HeroCreateRequest
	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"JSON_ParseError": err.Error()})
		return
	}
	newHero, err := hh.usecase.CreateHero(token, req.HeroName, req.Potion)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"userId":    newHero.UserID,
		"hero_name": newHero.HeroName,
	})
}
