package models

import (
	"api/internal/security"
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

type Usuario struct {
	ID       uint64    `json:"id,omitempty"`
	Nome     string    `json:"nome,omitempty"`
	Nick     string    `json:"nick,omitempty"`
	Email    string    `json:"email,omitempty"`
	Senha    string    `json:"senha,omitempty"`
	CriadoEm time.Time `json:"criadoEm,omitempty"`
}

// Format validates recived data user for non-blank filds and spaces.
func (u *Usuario) Format(step string) error {
	if err := u.validation(step); err != nil {
		return err
	}

	if err := u.trim(step); err != nil {
		return err
	}
	return nil
}

func (u *Usuario) validation(step string) error {
	if u.Nome == "" {
		return errors.New("campo nome é obrigatório")
	}

	if u.Nick == "" {
		return errors.New("campo nick é obrigatório")
	}

	if u.Email == "" {
		return errors.New("campo e-mail é obrigatório")
	}

	if err := checkmail.ValidateFormat(u.Email); err != nil {
		return errors.New("o e-mail inserido é inválido")
	}

	if step == "signUp" && u.Senha == "" {
		return errors.New("campo senha é obrigatório")
	}
	return nil
}

func (u *Usuario) trim(step string) error {

	u.Nome = strings.TrimSpace(u.Nome)
	u.Nick = strings.TrimSpace(u.Nick)
	u.Email = strings.TrimSpace(u.Email)

	if step == "signUp" {
		pwHash, err := security.Hash(u.Senha)
		if err != nil {
			return err
		}
		u.Senha = string(pwHash)
	}
	return nil
}
