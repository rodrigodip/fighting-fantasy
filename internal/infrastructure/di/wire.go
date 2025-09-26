package dependecy

import (
	"github.com/rodrigodip/fighting-fantasy/internal/application/auth"
	"github.com/rodrigodip/fighting-fantasy/internal/application/user"
	"github.com/rodrigodip/fighting-fantasy/internal/domain/auth"
	"github.com/rodrigodip/fighting-fantasy/internal/domain/user"
	"github.com/rodrigodip/fighting-fantasy/internal/infrastructure/config"
	"github.com/rodrigodip/fighting-fantasy/internal/infrastructure/persistence/mongodb/user"
	"github.com/rodrigodip/fighting-fantasy/internal/interface/http_handler/auth"
	"github.com/rodrigodip/fighting-fantasy/internal/interface/http_handler/user"
	"github.com/rodrigodip/fighting-fantasy/internal/pkg/security"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Container struct {
	User userhandler.UserHandlerRepo
	Auth authhandler.AuthHandlerRepo
}

func NewDependecyContainer(db mongo.Database) *Container {
	cfg := config.LoadConfig()

	UserRepo := mongodb.NewUserRepository(db.Collection("user"))
	userService := user.NewUserService(UserRepo)
	userUsecase := userapp.NewUserUseCase(userService)
	userHandler := userhandler.NewUserHandler(userUsecase)

	jwtService := security.NewJWTService(cfg.JWTSecret, cfg.JWTIssuer)
	authService := auth.NewAuthService(UserRepo, userService, jwtService)
	authUsecase := authapp.NewAuthUseCase(authService)
	authHandler := authhandler.NewAuthHandler(authUsecase)

	return &Container{
		User: userHandler,
		Auth: authHandler,
	}
}
