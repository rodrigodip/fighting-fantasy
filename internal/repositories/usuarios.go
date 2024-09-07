package repositories

import (
	"api/internal/models"
	"database/sql"
	"fmt"
	"log"
)

type usuarios struct {
	db *sql.DB
}

// NewUserRepo create a user repository
func NewUserRepo(db *sql.DB) *usuarios {
	return &usuarios{db}
}

// Create creates a user on DB
func (repo usuarios) Create(user models.Usuario) (uint64, error) {
	statement, err := repo.db.Prepare(
		"INSERT INTO usuarios (nome, nick, email, senha) VALUES (?,?,?,?)",
	)
	if err != nil {
		return 0, nil
	}
	defer statement.Close()

	result, err := statement.Exec(user.Nome, user.Nick, user.Email, user.Senha)
	if err != nil {
		return 0, nil
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		return 0, nil
	}

	return uint64(lastId), nil
}

// FindAll finds all users on DB
func (repo usuarios) FindAll(nameNikeFilter string) ([]models.Usuario, error) {

	nameNikeFilter = fmt.Sprintf("%%%s%%", nameNikeFilter)

	tuples, err := repo.db.Query(
		"SELECT id, nome, nick, email, criadoEm FROM usuarios WHERE nome LIKE ? or nick LIKE ?",
		nameNikeFilter, nameNikeFilter)
	if err != nil {
		log.Fatal("Erro recemendo usuarios do BD")
	}
	defer tuples.Close()

	var allUsers []models.Usuario

	for tuples.Next() {
		var usuario models.Usuario
		if err := tuples.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm); err != nil {
			return nil, err
		}

		allUsers = append(allUsers, usuario)
	}

	return allUsers, nil
}

// FindOne finds a user by id
func (repo usuarios) FindOne(userId uint64) (models.Usuario, error) {

	user, err := repo.db.Query(
		"select id, nome, nick, email, criadoEm from usuarios where id = ?", userId)
	if err != nil {
		return models.Usuario{}, err
	}
	defer user.Close()

	var foundUser models.Usuario
	if user.Next() {
		if err := user.Scan(
			&foundUser.ID,
			&foundUser.Nome,
			&foundUser.Nick,
			&foundUser.Email,
			&foundUser.CriadoEm); err != nil {
			return models.Usuario{}, err
		}
	}

	return foundUser, nil
}
