package heroapp

import "github.com/rodrigodip/fighting-fantasy/internal/domain/hero"

type Repository interface {
	FindByOwner(userID string) (*hero.Hero, error)
	//RegisterHero(hero hero.Hero) error
}
