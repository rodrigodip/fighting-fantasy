package webhandlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rodrigodip/fighting-fantasy/internal/application/hero"
	"github.com/rodrigodip/fighting-fantasy/internal/interface/web/viewmodels"
)

type HeroWebHandler struct {
	heroUseCase *heroapp.HeroUseCase
}

func NewHeroWebHandler(heroUseCase *heroapp.HeroUseCase) *HeroWebHandler {
	return &HeroWebHandler{
		heroUseCase: heroUseCase,
	}
}

func (h *HeroWebHandler) Detail(c *gin.Context) {
	// userID := c.Param("userId")
	// userID := "3035195610914029568"
	foundHero, err := h.heroUseCase.FindByUser("7971340423962820608")
	log.Printf("webHandler.Details(): %v", foundHero.Stats.InitialHP)
	if err != nil {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{
			"error": err.Error(),
		})
		return
	}
	viewModel := viewmodels.NewHeroViewModel(foundHero)
	c.HTML(http.StatusOK, "index.html", gin.H{
		"Title":     "My Hero",
		"HeroName":  viewModel.HeroName,
		"InitialHP": viewModel.InitialHP,
		"CurrentHP": viewModel.CurrentHP,
	})
}
