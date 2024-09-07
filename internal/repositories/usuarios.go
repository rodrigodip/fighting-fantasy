package repositories

import (
	"api/internal/models"
	"database/sql"
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
