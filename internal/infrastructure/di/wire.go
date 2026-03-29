package dependecy

import (
	"github.com/gorilla/sessions"
	"github.com/rodrigodip/fighting-fantasy/internal/application/hero"
	"github.com/rodrigodip/fighting-fantasy/internal/application/user"
	"github.com/rodrigodip/fighting-fantasy/internal/domain/hero"
	"github.com/rodrigodip/fighting-fantasy/internal/domain/user"
	heroRepository "github.com/rodrigodip/fighting-fantasy/internal/infrastructure/persistence/mongodb/hero"
	userRepository "github.com/rodrigodip/fighting-fantasy/internal/infrastructure/persistence/mongodb/user"
	"github.com/rodrigodip/fighting-fantasy/internal/interface/api/hero"
	"github.com/rodrigodip/fighting-fantasy/internal/interface/api/user"
	webmiddleware "github.com/rodrigodip/fighting-fantasy/internal/interface/web"
	webhandlers "github.com/rodrigodip/fighting-fantasy/internal/interface/web/handlers"
	"github.com/rodrigodip/fighting-fantasy/internal/pkg/security"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Container struct {
	ApiUserHandlers *userhandler.UserHandler
	ApiHeroHandlers *herohandler.HeroHandler

	WebAuthMiddleware *webmiddleware.WebAuthMiddleware
	WebUserHandlers   *webhandlers.UserWebHandler
	WebPagesHandlers  *webhandlers.PagesWebHandler
}

// TODO: Must decouple secret and issuer from Dependency constructor
func NewDependecyContainer(db mongo.Database, sessionStore sessions.Store, secret, issuer string) *Container {
	// Aplication Dependecies
	userRepo := userRepository.NewMongoUserRepository(db.Collection("user"))
	userService := usr.NewUserService()
	userAuth := security.NewJWTService(secret, issuer)
	userUsecase := userapp.NewusrUseCase(userService, userRepo, userAuth)

	heroRepo := heroRepository.NewMongoHeroRepository(db.Collection("hero"))
	heroService := hero.NewHeroService()
	heroUsecase := heroapp.NewHeroUseCase(heroService, heroRepo)

	// API Dependencies
	userApiHandler := userhandler.NewUserHandler(userUsecase)
	heroApiHandler := herohandler.NewHeroHandler(heroUsecase)

	// Web Dependecies
	webAuth := webmiddleware.NewWebAuthMiddleware(sessionStore, userAuth)
	userWebHandler := webhandlers.NewUserWebHandler(userUsecase, sessionStore)
	pagesWebHandler := webhandlers.NewPagesWebHandler(heroUsecase, userUsecase, sessionStore)

	// &Container Serves API and Web handlers to Routes
	return &Container{
		ApiUserHandlers: userApiHandler,
		ApiHeroHandlers: heroApiHandler,

		WebAuthMiddleware: webAuth,
		WebUserHandlers:   userWebHandler,
		WebPagesHandlers:  pagesWebHandler,
	}
}
