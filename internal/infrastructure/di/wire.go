package dependecy

import (
	"github.com/rodrigodip/fighting-fantasy/internal/application/hero"
	"github.com/rodrigodip/fighting-fantasy/internal/application/user"
	"github.com/rodrigodip/fighting-fantasy/internal/domain/hero"
	"github.com/rodrigodip/fighting-fantasy/internal/domain/user"
	heroRepository "github.com/rodrigodip/fighting-fantasy/internal/infrastructure/persistence/mongodb/hero"
	userRepository "github.com/rodrigodip/fighting-fantasy/internal/infrastructure/persistence/mongodb/user"
	"github.com/rodrigodip/fighting-fantasy/internal/interface/http_handler/hero"
	"github.com/rodrigodip/fighting-fantasy/internal/interface/http_handler/user"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Container struct {
	User *userhandler.UserHandler
	Hero *herohandler.HeroHandler
}

func NewDependecyContainer(db mongo.Database) *Container {

	userRepo := userRepository.NewMongoUserRepository(db.Collection("user"))
	userService := usr.NewUserService(userRepo)
	userUsecase := userapp.NewusrUseCase(userService)
	userHandler := userhandler.NewUserHandler(userUsecase)

	heroRepo := heroRepository.NewMongoHeroRepository(db.Collection("hero"))
	heroService := hero.NewHeroService(heroRepo)
	heroUsecase := heroapp.NewHeroUseCase(heroService)
	heroHandler := herohandler.NewHeroHandler(heroUsecase)
	return &Container{
		User: userHandler,
		Hero: heroHandler,
	}
}
