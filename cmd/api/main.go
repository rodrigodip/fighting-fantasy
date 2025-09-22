package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/rodrigodip/fighting-fantasy/internal/application/auth"
	"github.com/rodrigodip/fighting-fantasy/internal/application/user"
	"github.com/rodrigodip/fighting-fantasy/internal/domain/auth"
	"github.com/rodrigodip/fighting-fantasy/internal/domain/user"
	"github.com/rodrigodip/fighting-fantasy/internal/infrastructure/config"
	"github.com/rodrigodip/fighting-fantasy/internal/infrastructure/http/gin"
	"github.com/rodrigodip/fighting-fantasy/internal/infrastructure/persistence/mongodb/user"
	"github.com/rodrigodip/fighting-fantasy/internal/interface/http_handler/auth"
	"github.com/rodrigodip/fighting-fantasy/internal/interface/http_handler/user"
	"github.com/rodrigodip/fighting-fantasy/internal/pkg/security"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	cfg := config.LoadConfig()

	db, err := config.NewMongoDBConnection(context.Background())
	if err != nil {
		log.Fatalf("Error trying to connect to database, error=%s \n", err.Error())
		return
	}

	//TODO: REFACTOR: Implement a dependency initialization func
	userRepo := mongodb.NewUserRepository(db.Collection("user"))
	userService := user.NewUserService(userRepo)
	userUsecase := userapp.NewUserUseCase(userService)
	userHandler := userhandler.NewUserHandler(userUsecase)

	jwtService := security.NewJWTService(cfg.JWTSecret, cfg.JWTIssuer)
	authService := auth.NewAuthService(userRepo, userService, jwtService)
	authUsecase := authapp.NewAuthUseCase(authService)
	authHandler := authhandler.NewAuthHandler(authUsecase)

	router := gin.Default()
	routes.InitUserGroup(&router.RouterGroup, userHandler)
	routes.InitAuthGroup(&router.RouterGroup, authHandler)
	//TODO: DELETE: This is a protected route for email verification
	//	router.GET("/verify", interfaces.AuthMiddleware(cfg.JWTSecret, auth.RoleUser), authHandler.VerifyEmail)
	err = router.Run(cfg.HTTPPort)
	if err != nil {
		log.Fatal(err)
	}
}
