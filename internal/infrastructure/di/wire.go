package dependecy

import (
	"github.com/rodrigodip/fighting-fantasy/internal/application/auth"
	"github.com/rodrigodip/fighting-fantasy/internal/application/hero"
	"github.com/rodrigodip/fighting-fantasy/internal/application/user"
	"github.com/rodrigodip/fighting-fantasy/internal/domain/auth"
	"github.com/rodrigodip/fighting-fantasy/internal/domain/hero"
	"github.com/rodrigodip/fighting-fantasy/internal/domain/user"
	"github.com/rodrigodip/fighting-fantasy/internal/infrastructure/config"
	"github.com/rodrigodip/fighting-fantasy/internal/infrastructure/persistence/mongodb"
	"github.com/rodrigodip/fighting-fantasy/internal/interface/http_handler/auth"
	"github.com/rodrigodip/fighting-fantasy/internal/interface/http_handler/hero"
	"github.com/rodrigodip/fighting-fantasy/internal/interface/http_handler/user"
	"github.com/rodrigodip/fighting-fantasy/internal/pkg/security"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Container struct {
	User userhandler.UserHandlerRepo
	Auth authhandler.AuthHandlerRepo
	Hero herohandler.HeroHandlerRepo
}

func NewDependecyContainer(db mongo.Database) *Container {
	cfg := config.LoadConfig()

	userRepo := mongodb.NewMongoRepository(db.Collection("user"))
	userService := user.NewUserService(userRepo)
	userUsecase := userapp.NewUserUseCase(userService)
	userHandler := userhandler.NewUserHandler(userUsecase)

	jwtService := security.NewJWTService(cfg.JWTSecret, cfg.JWTIssuer)
	authService := auth.NewAuthService(userRepo, userService, jwtService)
	authUsecase := authapp.NewAuthUseCase(authService)
	authHandler := authhandler.NewAuthHandler(authUsecase)

	heroRepo := mongodb.NewMongoRepository(db.Collection("hero"))
	heroService := hero.NewHeroService(heroRepo)
	heroUsecase := heroapp.NewHeroUseCase(heroService)
	heroHandler := herohandler.NewHeroHandler(heroUsecase)

	return &Container{
		User: userHandler,
		Auth: authHandler,
		Hero: heroHandler,
	}
}
