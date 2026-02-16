package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rodrigodip/fighting-fantasy/internal/application/hero"
	"github.com/rodrigodip/fighting-fantasy/internal/interface/web/templates"
	"github.com/rodrigodip/fighting-fantasy/internal/interface/web/viewmodels"
)

type HeroHandler struct {
	heroUseCase *heroapp.HeroUseCase
	renderer    *templates.Renderer
}

func NewHeroHandler(heroUseCase *heroapp.HeroUseCase, renderer *templates.Renderer) *HeroHandler {
	return &HeroHandler{
		heroUseCase: heroUseCase,
		renderer:    renderer,
	}
}

func (h *HeroHandler) Detail(c *gin.Context) {
	userID := c.Param("userId")
	foundHero, err := h.heroUseCase.FindByUser(userID)
	if err != nil {
		c.HTML(http.StatusNotFound, "error.html", gin.H{
			"error": "Hero not found",
		})
		return
	}
	viewModel := viewmodels.NewHeroViewModel(foundHero)
	h.renderer.Render(c, http.StatusOK, "hero/detail.html", gin.H{
		"HeroName": viewModel.HeroName,
		"HP":       viewModel.HP,
	})
}
