package userapp

import "github.com/rodrigodip/fighting-fantasy/internal/domain/user"

type UserRepository interface {
	RegisterUser(u usr.User) error
	FindByEmail(email string) (*usr.User, error)
	FindById(id string) (*usr.User, error)
	Update(id, field, newData string) error

	//TODO: ResendVeriyEmail()
	//FindHero(userID string) (*Hero, error)
}

type Authentication interface {
	GenerateToken(userID, role string) (string, error)
	ValidateToken(token string) (string, error)
	SendVerifyEmail(userID, name, email, role string) error
}
