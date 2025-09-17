package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/rodrigodip/fighting-fantasy/internal/application/auth"
	"github.com/rodrigodip/fighting-fantasy/internal/application/user"
	"github.com/rodrigodip/fighting-fantasy/internal/domain/user"
	"github.com/rodrigodip/fighting-fantasy/internal/infrastructure/config"
	"github.com/rodrigodip/fighting-fantasy/internal/infrastructure/http/gin"
	"github.com/rodrigodip/fighting-fantasy/internal/infrastructure/persistence/mongodb/user"
	"github.com/rodrigodip/fighting-fantasy/internal/interface/http_handler/auth"
	"github.com/rodrigodip/fighting-fantasy/internal/interface/http_handler/user"
	"github.com/rodrigodip/fighting-fantasy/internal/pkg/security"
)

func main() {
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

	//TODO: REFACTOR: Create a Config func to load SECRETS
	jwtService := security.NewJWTService("Cjy+8nlgKa7FJUsR8pEX0eV0l/Nu6pI0yYBaaXYgL0uE3cczRHpOqkAeqT9vPbGf", "ffantasy-app")
	authUsecase := authapp.NewAuthUseCase(userRepo, userService, jwtService)
	authHandler := authhandler.NewAuthHandler(authUsecase)

	router := gin.Default()
	routes.InitUserGroup(&router.RouterGroup, userHandler)
	routes.InitAuthGroup(&router.RouterGroup, authHandler)
	err = router.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
