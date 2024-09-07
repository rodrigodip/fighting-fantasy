package main

import (
	"api/internal/database/config"
	"api/internal/server"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.LoadEnv()

	port := fmt.Sprintf(":%d", config.ApiAddr)

	fmt.Printf("Rodando API na porta: %s \n", port)

	fmt.Printf("String de conexão: %s", config.StrinConnect)

	r := server.Up()
	log.Fatal(http.ListenAndServe(port, r))
}
