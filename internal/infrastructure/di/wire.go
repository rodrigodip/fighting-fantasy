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
	"github.com/rodrigodip/fighting-fantasy/internal/pkg/security"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Container struct {
	User *userhandler.UserHandler
	Hero *herohandler.HeroHandler
}

// TODO: Must decouple secret and issuer from Dependency constructor
func NewDependecyContainer(db mongo.Database, secret, issuer string) *Container {

	userRepo := userRepository.NewMongoUserRepository(db.Collection("user"))
	userService := usr.NewUserService()
	userAuth := security.NewJWTService(secret, issuer)
	userUsecase := userapp.NewusrUseCase(userService, userRepo, userAuth)
	userHandler := userhandler.NewUserHandler(userUsecase)

	heroRepo := heroRepository.NewMongoHeroRepository(db.Collection("hero"))
	heroService := hero.NewHeroService()
	heroUsecase := heroapp.NewHeroUseCase(heroService, heroRepo)
	heroHandler := herohandler.NewHeroHandler(heroUsecase)
	return &Container{
		User: userHandler,
		Hero: heroHandler,
	}
}
