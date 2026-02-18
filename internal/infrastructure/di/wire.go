package dependecy

import (
	"github.com/rodrigodip/fighting-fantasy/internal/application/hero"
	"github.com/rodrigodip/fighting-fantasy/internal/application/user"
	"github.com/rodrigodip/fighting-fantasy/internal/domain/hero"
	"github.com/rodrigodip/fighting-fantasy/internal/domain/user"
	heroRepository "github.com/rodrigodip/fighting-fantasy/internal/infrastructure/persistence/mongodb/hero"
	userRepository "github.com/rodrigodip/fighting-fantasy/internal/infrastructure/persistence/mongodb/user"
	"github.com/rodrigodip/fighting-fantasy/internal/interface/api/hero"
	"github.com/rodrigodip/fighting-fantasy/internal/interface/api/user"
	webhandlers "github.com/rodrigodip/fighting-fantasy/internal/interface/web/handlers"
	"github.com/rodrigodip/fighting-fantasy/internal/pkg/security"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Container struct {
	ApiUserHandlers *userhandler.UserHandler
	ApiHeroHandlers *herohandler.HeroHandler

	WebUserHandlers *webhandlers.UserWebHandler
	WebHeroHandlers *webhandlers.HeroWebHandler
}

// TODO: Must decouple secret and issuer from Dependency constructor
func NewDependecyContainer(db mongo.Database, secret, issuer string) *Container {

	// API Dependecies
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
	userWebHandler := webhandlers.NewUserWebHandler(userUsecase)
	heroWebHandler := webhandlers.NewHeroWebHandler(heroUsecase)

	// &Container Serves API and Web handlers to Routes
	return &Container{
		ApiUserHandlers: userApiHandler,
		ApiHeroHandlers: heroApiHandler,

		WebUserHandlers: userWebHandler,
		WebHeroHandlers: heroWebHandler,
	}
}
