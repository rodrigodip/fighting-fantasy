package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
	"github.com/rodrigodip/fighting-fantasy/internal/infrastructure/config"
	"github.com/rodrigodip/fighting-fantasy/internal/infrastructure/di"
	"github.com/rodrigodip/fighting-fantasy/internal/infrastructure/http/gin"
)

// @title Figthing Fantasy
// @version 1.0
// @description A solo Game-Book RPG
// @host localhost:8080
// @BasePath /
// @schemes http
// @license MIT
func main() {
	_ = godotenv.Load()
	cfg := config.LoadConfig()

	db, err := config.NewMongoDBConnection(context.Background())
	if err != nil {
		log.Fatalf("Error trying to connect to database, error=%s \n", err.Error())
		return
	}
	// Create cookie store with secret key
	sessionStore := sessions.NewCookieStore([]byte(cfg.SessionSecret))
	sessionStore.Options = cfg.SessionOpt

	di := dependecy.NewDependecyContainer(*db, sessionStore, cfg.JWTSecret, cfg.JWTIssuer)

	router := gin.Default()
	router.LoadHTMLGlob(cfg.TemplatesPath)
	router.Static("/static", cfg.StaticPath)

	routes.InitUserGroup(&router.RouterGroup, di.ApiUserHandlers, *cfg)
	routes.InitHeroGroup(&router.RouterGroup, di.ApiHeroHandlers, *cfg)

	routes.InitWebAuthGroup(&router.RouterGroup, di)
	routes.InitWebPagesGroup(&router.RouterGroup, di)

	err = router.Run(cfg.HTTPPort)
	if err != nil {
		log.Fatal(err)
	}
}
