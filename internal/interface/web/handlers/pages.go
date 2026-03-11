package webhandlers

import (
	"html/template"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"github.com/rodrigodip/fighting-fantasy/internal/application/hero"
	"github.com/rodrigodip/fighting-fantasy/internal/interface/web/viewmodels"
)

type PagesWebHandlerRepo interface {
	AuthPageHandler(c *gin.Context)
	DashboardPageHandler(c *gin.Context)
	AdventurePageHandlerView(c *gin.Context)
	GameOverPageHandler(c *gin.Context)
}

type PagesWebHandler struct {
	usecase      *heroapp.HeroUseCase
	sessionStore sessions.Store
}

func NewPagesWebHandler(uc *heroapp.HeroUseCase, store sessions.Store) *PagesWebHandler {
	return &PagesWebHandler{
		usecase:      uc,
		sessionStore: store,
	}
}

// AuthPageHandler renders the authentication page
// GET /
func (uc *PagesWebHandler) AuthPageHandler(c *gin.Context) {
	// TODO: Render auth.html template
	// 1. Check if user is already logged in
	// 2. If logged in, redirect to /dashboard
	// 3. Otherwise, render login/signup page

	// Check if the user is logged in
	// session, _ := uc.sessionStore.Get(c.Request, "user-session")
	// if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
	// 	c.Header("HX-Redirect", "/dashboard")
	// 	c.Status(http.StatusOK)
	// 	return
	// }

	tmpl := template.Must(template.New("").Funcs(GetTemplateFunctions()).ParseFiles(
		"internal/interface/web/templates/layouts/base.html",
		"internal/interface/web/templates/pages/auth.html",
	))
	data := viewmodels.PageData{
		Title: "Login",
	}
	tmpl.ExecuteTemplate(c.Writer, "base.html", data)
}

// DashboardPageHandler renders the dashboard page
// GET /dashboard
func (uc *PagesWebHandler) DashboardPageHandler(c *gin.Context) {
	// TODO: Render dashboard.html template
	// 1. Get user from session (or redirect to / if not logged in)
	// 2. Fetch user's hero (if exists)
	// 3. Render dashboard with user and hero data

	// Example rendering:
	// tmpl := template.Must(template.ParseFiles(
	// 	"internal/interface/web/templates/layouts/base.html",
	// 	"internal/interface/web/templates/pages/dashboard.html",
	// ))
	//
	// data := PageData{
	// 	Title: "Dashboard",
	// 	User:  currentUser,
	// 	Hero:  userHero, // nil if no hero created yet
	// }
	//
	// tmpl.ExecuteTemplate(w, "base.html", data)
}

// AdventurePageHandlerView renders the adventure page
// GET /adventure
func (uc *PagesWebHandler) AdventurePageHandlerView(c *gin.Context) {
	// TODO: Render adventure.html template
	// 1. Get user from session (or redirect to / if not logged in)
	// 2. Get user's hero (redirect to /dashboard if no hero)
	// 3. Get current adventure page (default to page 1)
	// 4. Fetch story content for current page
	// 5. Render adventure template with hero stats and story

	// Example rendering:
	// tmpl := template.Must(template.ParseFiles(
	// 	"internal/interface/web/templates/layouts/base.html",
	// 	"internal/interface/web/templates/pages/adventure.html",
	// 	"internal/interface/web/templates/partials/dice.html",
	// ))
	//
	// data := PageData{
	// 	Title:       "Adventure",
	// 	User:        currentUser,
	// 	Hero:        userHero,
	// 	CurrentPage: currentPageNumber,
	// 	Story:       loreSection,
	// }
	//
	// tmpl.ExecuteTemplate(w, "base.html", data)
}

// GameOverPageHandler renders the game over page
// GET /gameover
func (uc *PagesWebHandler) GameOverPageHandler(c *gin.Context) {
	// TODO: Render gameover.html template
	// 1. Get user from session
	// 2. Calculate or fetch final statistics
	// 3. Render game over page with stats

	// Example rendering:
	// tmpl := template.Must(template.ParseFiles(
	// 	"internal/interface/web/templates/layouts/base.html",
	// 	"internal/interface/web/templates/pages/gameover.html",
	// ))
	//
	// data := PageData{
	// 	Title: "Game Over",
	// 	User:  currentUser,
	// 	Stats: &GameStats{
	// 		BattlesWon:    12,
	// 		GoldCollected: 45,
	// 		AreasExplored: 8,
	// 	},
	// }
	//
	// tmpl.ExecuteTemplate(w, "base.html", data)
}

// GetTemplateFunctions returns custom template functions
func GetTemplateFunctions() template.FuncMap {
	return template.FuncMap{
		// dict creates a map from key-value pairs
		// Usage: {{template "partial" (dict "key1" "value1" "key2" "value2")}}
		"dict": func(values ...interface{}) (map[string]interface{}, error) {
			if len(values)%2 != 0 {
				return nil, nil
			}
			dict := make(map[string]interface{}, len(values)/2)
			for i := 0; i < len(values); i += 2 {
				key, ok := values[i].(string)
				if !ok {
					return nil, nil
				}
				dict[key] = values[i+1]
			}
			return dict, nil
		},

		// multiply multiplies two numbers
		// Usage: {{multiply .Hero.Strength 5}}
		"multiply": func(a, b int) int {
			return a * b
		},

		// divide divides two numbers
		// Usage: {{divide .Hero.CurrentStrength .Hero.Strength}}
		"divide": func(a, b int) int {
			if b == 0 {
				return 0
			}
			return a / b
		},
	}
}
