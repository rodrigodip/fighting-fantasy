package main

import (
	"api/internal/config"
	"api/internal/server"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.LoadEnv()

	port := fmt.Sprintf(":%d", config.ApiAddr)

	fmt.Printf("Rodando API na porta%s \n", port)

	r := server.Up()
	log.Fatal(http.ListenAndServe(port, r))
}
