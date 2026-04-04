package webhandlers

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"github.com/rodrigodip/fighting-fantasy/internal/application/hero"
	"github.com/rodrigodip/fighting-fantasy/internal/interface/web/viewmodels"
)

type HeroWebHandlerRepo interface {
	HeroCreateHandler(c *gin.Context)
}

type HeroWebHandler struct {
	usecase      *heroapp.HeroUseCase
	sessionStore sessions.Store
}

func NewHeroWebHandler(uc *heroapp.HeroUseCase, store sessions.Store) *HeroWebHandler {
	return &HeroWebHandler{
		usecase:      uc,
		sessionStore: store,
	}
}
func (uc *HeroWebHandler) HeroCreateHandler(c *gin.Context) {
	log.Printf("context keys: %v", c.Keys)
	userID, exists := c.Get("user_id")
	if !exists {
		log.Println("User not found")
		// c.Redirect(303, "/")
		return
	}
	var req viewmodels.HeroCreateReq
	if err := c.ShouldBind(&req); err != nil {
		// TODO: Create a Error Container on Dashboard.html
		log.Println("Erro Binding req")
		return
	}
	_, err := uc.usecase.CreateHero(userID.(string), req.HeroName, req.Potion)
	if err != nil {
		log.Printf("%v", err)
		return
	}
	// Return the re-rendered template fragment for HTMX to swap
	c.Header("HX-Redirect", "/dashboard")
	c.Status(200)
}
