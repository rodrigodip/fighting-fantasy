package heroapp

import "github.com/rodrigodip/fighting-fantasy/internal/domain/hero"

type Repository interface {
	RegisterHero(hero hero.Hero) error
}
