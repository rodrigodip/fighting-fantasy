package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/rodrigodip/fighting-fantasy/internal/application/user"
	"github.com/rodrigodip/fighting-fantasy/internal/domain/user"
	"github.com/rodrigodip/fighting-fantasy/internal/infrastructure/config"
	"github.com/rodrigodip/fighting-fantasy/internal/infrastructure/persistence/mongodb"
	"github.com/rodrigodip/fighting-fantasy/internal/interface/http_handler/user"
)

func main() {
	db, err := mongoconfig.NewMongoDBConnection(context.Background())
	if err != nil {
		log.Fatalf("Error trying to connect to database, error=%s \n", err.Error())
		return
	}
	defer func() {
		if err := db.Client().Disconnect(context.TODO()); err != nil {
			log.Fatalf("Error trying to disconnect from database, error=%s \n", err.Error())
			return
		}
	}()
	//TODO: REFACTOR: Implement a dependency initialization func
	userRepo := mongodb.NewUserRepository(db.Collection("user"))
	userService := user.NewUserService(userRepo)
	userUsecase := userapp.NewUserUseCase(userService)
	userHandler := userhandler.NewUserHandler(userUsecase)

	//TODO: REFACTOR: Implement a router
	r := gin.Default()
	r.POST("users/", userHandler.CreateUser)

	r.Run(":8080")
}
