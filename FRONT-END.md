# interface/web

```
/<handlers>     # Golang files to handle web requisitions. This files must call <application> usecases.
/<static>       # Utils scripst like vanila javascript and CSS
/<templates>    # HTML files
    /<layouts>  # Base Layout
    /<pages>
    /<partials>
/<viewmodes>    # Data Transfer Layers to handle front end data flow.
```

## Features

### Authentication (`auth.html`)
- Glassmorphic login/signup forms
- Tab-based interface with Shield (Login) and Sword (Signup) icons
- HTMX-powered form submission
- Atmospheric forest illustration background

### Dashboard (`dashboard.html`)
- **Create Hero**: Name input and potion selection (Dexterity, Strength, Fortune)
- **Hero Status**: View stats, equipped items, and backpack inventory
- **Account Settings**: Manage account and delete hero

### Adventure (`adventure.html`)
- **Story System**: Narrative text with multiple choice decisions
- **Battle Mode**: Dice-based combat with visual feedback
- **Hero Stats Sidebar**: Real-time stat tracking
- **Inventory Display**: Quick view of provisions, gold, and jewels

### Game Over (`gameover.html`)
- Death screen with atmospheric visuals
- Final statistics (battles won, gold collected, areas explored)
- Options to return to dashboard or retry

## Backend Implementation Guide

All Go handlers are located in `/internal/interface/web/handlers/`. Each handler file contains:

### Auth Handlers (`auth.go`)
- `AuthLoginHandler` - POST /auth/login
- `AuthSignUpHandler` - POST /auth/signup
- `AuthLogoutHandler` - POST /auth/logout
- Auth Handlers must use `userapp` package functions

### Hero Handlers (`hero.go`)
- `HeroCreateHandler` - POST /hero/create
- `HeroStatusHandler` - GET /hero/status
- `HeroDeleteHandler` - DELETE /hero
- `HeroUpdateStatsHandler` - PUT /hero/stats
- Hero Handlers must use `heroapp` package functions

### Adventure Handlers (`adventure.go`)
- `AdventurePageHandler` - GET /adventure/page/:pageId
- `BattleAttackHandler` - POST /adventure/battle/attack
- `HeroStatsHandler` - GET /adventure/hero-stats
- `AdventureStartHandler` - POST /adventure/start
- `AdventureEndHandler` - POST /adventure/end
- Adventure Handlers must use `advapp` package functions

### Page Handlers (`pages.go`)
- `AuthPageHandler` - GET /
- `DashboardPageHandler` - GET /dashboard
- `AdventurePageHandlerView` - GET /adventure
- `GameOverPageHandler` - GET /gameover

**Missing Aplication Use Case:**
If there isn't any package to handle a specific use case, generate it.

**Investigate First** Before write any code read existent use cases packages and write similar estructure and code.

## A use case example
```
// <-- /application/hero/usecase.go -->
package heroapp

import (
	"fmt"
	"log"

	"github.com/rodrigodip/fighting-fantasy/internal/domain/hero"
)

type HeroUseCase struct {
	service *hero.Service
	repo    Repository
}

func NewHeroUseCase(s *hero.Service, r Repository) *HeroUseCase {
	return &HeroUseCase{service: s, repo: r}
}
// CreateHero use case, creates a hero on mongoDB database
func (uc *HeroUseCase) CreateHero(userID, heroName, potion string) (*hero.Hero, error) {
	foundHero, _ := uc.repo.FindByOwner(userID)
	if err := uc.service.HasHero(userID, *foundHero); err != nil {
		return &hero.Hero{}, fmt.Errorf("CreateHero: %v", err)
	}
	if err := uc.service.ValidateInput(
		heroName, potion,
	); err != nil {
		return &hero.Hero{}, fmt.Errorf("CreateHero: %v", err)
	}
	newHero := hero.Hero{
		UserID:   userID,
		HeroName: heroName,
		Stats: hero.Stats{
			InitialHP: 10,
			CurrentHP: 6,
		},
	}
	newHero, err := uc.service.SelectPotion(newHero, potion)
	if err != nil {
		return &hero.Hero{}, fmt.Errorf("CreateHero: %v", err)
	}
	if err := uc.repo.RegisterHero(newHero); err != nil {
		return &hero.Hero{}, fmt.Errorf("CreateHero: %v", err)
	}
	return &newHero, nil
}
```

## Routes examples
```
func InitWebAuthGroup(r *gin.RouterGroup, d *dependecy.Container) {
	r.GET("/auth/login", d.WebUserHandlers.AuthLoginHandler)
}
```
- Use gin groups to better routes organization
This is a exemple how routes are called on `main.go`
```
router := gin.Default()
routes.InitWebAuthGroup(&router.RouterGroup, di) // di is a dependency container
```
## HTMX Integration

The application uses HTMX for dynamic interactions without full page reloads:

- Form submissions with `hx-post`, `hx-get`
- Target updates with `hx-target` and `hx-swap`
- Confirmation dialogs with `hx-confirm`
- Smooth content transitions

## Vanilla JavaScript

All client-side interactivity is handled in `/static/js/app.js`:

- `switchTab(tabName)` - Tab navigation
- `selectPotion(type, element)` - Potion selection UI
- `rollDice(id, sides)` - Dice rolling animation
- `executeBattleRound()` - Battle calculation
- `addBattleLogEntry(message, type)` - Battle log updates

## Visual Design

### Color Palette
- **Background**: Deep forest green (#0a0e0d)
- **Primary**: Forest green (#4a7c59)
- **Gold Accents**: (#d4af37) for headings and highlights
- **Card Glass**: Translucent dark (#121c17 with 80% opacity)

### Glassmorphism Effects
All cards use:
- `backdrop-filter: blur(12px)`
- Semi-transparent backgrounds
- Subtle borders with low opacity
- Shadow layering for depth

### Typography
- **Headings**: Cinzel (serif)
- **Body**: Crimson Text (serif)
- Maintains classic gamebook aesthetic

## Running the Application

1. **Run Docker container**
    `make dev-build` - use if any docker related file was changed
    `make dev` - use for run docker compose up
    `make dev-down` - use for docker compose down
    `make clear` - use to remove containers volumes

2. **Access the application**:
   ```
   http://localhost:8080
   ```

## Development Notes

- All templates use Go's `html/template` package
- Backend handlers need to be implemented with application layer logic
- HTMX handles most AJAX interactions
- Tailwind CSS is loaded via CDN (can be optimized for production)
- Static assets served from `/static/` directory

## Template Data Structures

Templates must use DTO's to represent entities served from `/domain/`
