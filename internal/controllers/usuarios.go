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

	if err = user.Format("signUp"); err != nil {
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

	nameNickeFilter := strings.ToLower(r.URL.Query().Get("user"))

	db, err := database.Connection()
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	users := repositories.NewUserRepo(db)
	allUser, err := users.FindAll(nameNickeFilter)
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, allUser)
}

// BuscarUsuario requestes a user by ID
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

// AtualizarUsuario updates a user by ID
func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {

	useParams := mux.Vars(r)

	ID, err := strconv.ParseUint(useParams["userId"], 10, 64)
	if err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}

	bodyRequest, err := io.ReadAll(r.Body)
	if err != nil {
		responses.Err(w, http.StatusUnprocessableEntity, err)
		return
	}

	var newUser models.Usuario
	if err = json.Unmarshal(bodyRequest, &newUser); err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}

	if err = newUser.Format("edit"); err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connection()
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	user := repositories.NewUserRepo(db)
	if err = user.UpDate(ID, newUser); err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

// DeleteUsuario deletes a user by ID
func DeleteUsuario(w http.ResponseWriter, r *http.Request) {

	userParams := mux.Vars(r)

	ID, err := strconv.ParseUint(userParams["userId"], 10, 64)
	if err != nil {
		responses.Err(w, http.StatusBadRequest, err)
	}

	db, err := database.Connection()
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
	}
	defer db.Close()

	user := repositories.NewUserRepo(db)
	if err = user.Delete(ID); err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
	}

	responses.JSON(w, http.StatusNoContent, nil)

}
