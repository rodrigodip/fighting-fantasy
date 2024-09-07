package controllers

import (
	"api/internal/database"
	"api/internal/models"
	"api/internal/repositories"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// CriarUsuarios creates a user
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	bodyRequest, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	var user models.Usuario
	if err = json.Unmarshal(bodyRequest, &user); err != nil {
		log.Fatal(err)
	}

	db, err := database.Connection()
	if err != nil {
		log.Fatal(err)
	}

	newrRepo := repositories.NewUserRepo(db)
	userId, err := newrRepo.Create(user)
	if err != nil {
		log.Fatal(err)
	}

	w.Write([]byte(fmt.Sprintf("ID inserido: %d", userId)))

}

// BuscarUsuarios requests all users
func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando todos os Usuário!"))
}

// BuscarUsuario requestes a user
func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando um Usuário!"))
}

// AtualizarUsuario updates a user
func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Atualizando Usuário!"))
}

// Deletes a user
func DeleteUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deletendo Usuário!"))
}
