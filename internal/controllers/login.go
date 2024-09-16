package controllers

import (
	"api/internal/auth"
	"api/internal/database"
	"api/internal/models"
	"api/internal/repositories"
	"api/internal/responses"
	"api/internal/security"
	"encoding/json"
	"io"
	"net/http"
)

// Login authenticate a user
func Login(w http.ResponseWriter, r *http.Request) {
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

	db, err := database.Connection()
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	newRepo := repositories.NewUserRepo(db)
	userOnDB, err := newRepo.FindByEmail(user.Email)
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
	}

	if err = security.PwVerify(userOnDB.Senha, user.Senha); err != nil {
		responses.Err(w, http.StatusUnauthorized, err)
		return
	}

	token, _ := auth.NewToken(userOnDB.ID)

	w.Write([]byte(token))
}
