package webhandlers

import (
	"html/template"
	"net/http"
)

// PageData represents common data for all pages
type PageData struct {
	Title       string
	User        *User
	Hero        *Hero
	CurrentPage int
	Story       *LoreSection
	Stats       *GameStats
}

// User represents an authenticated user
type User struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
}

// GameStats represents final adventure statistics
type GameStats struct {
	BattlesWon    int `json:"battles_won"`
	GoldCollected int `json:"gold_collected"`
	AreasExplored int `json:"areas_explored"`
}

// AuthPageHandler renders the authentication page
// GET /
func AuthPageHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: Render auth.html template
	// 1. Check if user is already logged in
	// 2. If logged in, redirect to /dashboard
	// 3. Otherwise, render login/signup page

	tmpl := template.Must(template.ParseFiles(
		"internal/interface/web/templates/layouts/base.html",
		"internal/interface/web/templates/pages/auth.html",
	))

	data := PageData{
		Title: "Login",
	}

	tmpl.ExecuteTemplate(w, "base.html", data)
}

// DashboardPageHandler renders the dashboard page
// GET /dashboard
func DashboardPageHandler(w http.ResponseWriter, r *http.Request) {
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
func AdventurePageHandlerView(w http.ResponseWriter, r *http.Request) {
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
func GameOverPageHandler(w http.ResponseWriter, r *http.Request) {
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
