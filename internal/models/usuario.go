package models

import (
	"errors"
	"strings"
	"time"
)

type Usuario struct {
	ID       uint64    `json:"id,omitempty"`
	Nome     string    `json:"nome,omitempty"`
	Nick     string    `json:"nick,omitempty"`
	Email    string    `json:"email,omitempty"`
	Senha    string    `json:"senha,omitempty"`
	CriadoEm time.Time `json:"criadoEm,omitempty"`
}

// Format prepers the received user
func (u *Usuario) Format() error {
	if err := u.validation(); err != nil {
		return err
	}

	u.trim()
	return nil
}

func (u *Usuario) validation() error {
	if u.Nome == "" {
		return errors.New("campo nome é obrigatório")
	}
	if u.Nick == "" {
		return errors.New("campo nick é obrigatório")
	}
	if u.Email == "" {
		return errors.New("campo e-mail é obrigatório")
	}
	if u.Senha == "" {
		return errors.New("campo senha é obrigatório")
	}
	return nil
}

func (u *Usuario) trim() {
	u.Nome = strings.TrimSpace(u.Nome)
	u.Nick = strings.TrimSpace(u.Nick)
	u.Email = strings.TrimSpace(u.Email)
}
