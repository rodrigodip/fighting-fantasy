package webhandlers

import (
	"html/template"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"github.com/rodrigodip/fighting-fantasy/internal/application/hero"
	"github.com/rodrigodip/fighting-fantasy/internal/application/user"
	"github.com/rodrigodip/fighting-fantasy/internal/interface/web/viewmodels"
	weberrors "github.com/rodrigodip/fighting-fantasy/internal/pkg/errors/web_errors"
)

type HeroWebHandlerRepo interface {
	HeroCreateHandler(c *gin.Context)
}

type HeroWebHandler struct {
	usecase      *heroapp.HeroUseCase
	userusecase  *userapp.UserUseCase
	sessionStore sessions.Store
}

func NewHeroWebHandler(huc *heroapp.HeroUseCase, uuc *userapp.UserUseCase, store sessions.Store) *HeroWebHandler {
	return &HeroWebHandler{
		usecase:      huc,
		userusecase:  uuc,
		sessionStore: store,
	}
}

func (uc *HeroWebHandler) renderDashboardContent(c *gin.Context, feedback *viewmodels.Message) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.Redirect(303, "/")
		return
	}
	foundUser, err := uc.userusecase.FindById(userID.(string))
	if err != nil {
		c.Redirect(303, "/")
		return
	}
	currentUser := viewmodels.User{
		ID:    foundUser.UserID,
		Name:  foundUser.Name,
		Email: foundUser.Email,
	}
	var userHero *viewmodels.HeroViewModel
	foundHero, err := uc.usecase.FindByUser(foundUser.UserID)
	if err == nil {
		userHero = &viewmodels.HeroViewModel{
			Name:      foundHero.HeroName,
			Strength:  foundHero.Stats.InitialHP,
			Dexterity: foundHero.Stats.CurrentDex,
			Fortune:   foundHero.Stats.CurrentLuck,
		}
	}
	tmpl := template.Must(template.ParseFiles(
		"internal/interface/web/templates/pages/dashboard.html",
		"internal/interface/web/templates/partials/auth-feedback.html",
	))
	data := viewmodels.PageData{
		Title:    "Dashboard",
		User:     &currentUser,
		Hero:     userHero,
		FeedBack: feedback,
	}
	tmpl.ExecuteTemplate(c.Writer, "content", data)
}

func (uc *HeroWebHandler) HeroCreateHandler(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.Redirect(303, "/")
		return
	}
	var req viewmodels.HeroCreateReq
	if err := c.ShouldBind(&req); err != nil {
		uc.renderDashboardContent(c, &viewmodels.Message{Error: "Invalid hero name or potion selection"})
		return
	}
	_, err := uc.usecase.CreateHero(userID.(string), req.HeroName, req.Potion)
	if err != nil {
		uc.renderDashboardContent(c, &viewmodels.Message{Error: weberrors.ParseErrorForWeb(err.Error())})
		return
	}
	c.Header("HX-Redirect", "/dashboard")
	c.Status(200)
}
