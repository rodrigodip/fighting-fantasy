package controllers

import (
	"api/internal/database"
	"api/internal/models"
	"api/internal/repositories"
	"api/internal/responses"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

// CriarUsuarios creates a user
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	bodyRequest, err := io.ReadAll(r.Body)
	if err != nil {
		responses.Err(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.Usuario
	if err = json.Unmarshal(bodyRequest, &user); err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}

	if err = user.Format(); err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connection()
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	newrRepo := repositories.NewUserRepo(db)
	user.ID, err = newrRepo.Create(user)
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusCreated, user)
}

// BuscarUsuarios requests all users
func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	nameNikeFilter := strings.ToLower(r.URL.Query().Get("user"))

	db, err := database.Connection()
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	// trazer os usuarios do repositorio
	users := repositories.NewUserRepo(db)
	allUser, err := users.FindAll(nameNikeFilter)
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	// escrever o response para response
	responses.JSON(w, http.StatusOK, allUser)
}

// BuscarUsuario requestes a user
func BuscarUsuario(w http.ResponseWriter, r *http.Request) {

	userParams := mux.Vars(r)

	ID, err := strconv.ParseUint(userParams["userId"], 10, 64)
	if err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connection()
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	users := repositories.NewUserRepo(db)
	foundUser, err := users.FindOne(ID)
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, foundUser)
}

// AtualizarUsuario updates a user
func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Atualizando Usuário!"))
}

// Deletes a user
func DeleteUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deletendo Usuário!"))
}
