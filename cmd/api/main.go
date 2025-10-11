package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/rodrigodip/fighting-fantasy/internal/infrastructure/config"
	"github.com/rodrigodip/fighting-fantasy/internal/infrastructure/di"
	"github.com/rodrigodip/fighting-fantasy/internal/infrastructure/http/gin"
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

	di := dependecy.NewDependecyContainer(*db)

	router := gin.Default()
	routes.InitUserGroup(&router.RouterGroup, di.User)
	routes.InitHeroGroup(&router.RouterGroup, di.Hero, *cfg)

	err = router.Run(cfg.HTTPPort)
	if err != nil {
		log.Fatal(err)
	}
}
