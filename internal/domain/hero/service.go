package hero

import (
	"errors"
	"fmt"
	"os"

	"github.com/golang-jwt/jwt/v5"
	authErr "github.com/rodrigodip/fighting-fantasy/internal/pkg/errors"
	"github.com/rodrigodip/fighting-fantasy/internal/pkg/security"
)

type Service struct {
	service Repository
}

func NewHeroService(r Repository) *Service {
	return &Service{service: r}
}

// ValidateInput check input for incosistence
func (s *Service) ValidateInput(heroName, potion string) error {
	if heroName == "" {
		return errors.New("ValidadeHero: heroName is Required")
	}
	if len(heroName) < 3 {
		return errors.New("ValidadeHero: heroName must have more than 3 digits")
	}
	if len(heroName) > 20 {
		return errors.New("ValidadeHero: heroName must have less than 20 digits")
	}
	if potion == "" {
		return errors.New("ValidadeHero: a potion name is Required")
	}
	return nil
}

// Save saves a hero on DB
func (s *Service) Save(h Hero) error {
	if err := s.service.RegisterHero(h); err != nil {
		return fmt.Errorf("Save: %v", err)
	}
	return nil
}

// SelectPotion attach a potion to Hero
func (s *Service) SelectPotion(hero Hero, potion string) (Hero, error) {
	switch potion {
	case "dexterity":
		hero.Potions.Dexterity = true
		return hero, nil
	case "strength":
		hero.Potions.Strength = true
		return hero, nil
	case "fortune":
		hero.Potions.Fortune = true
		return hero, nil
	default:
		return Hero{}, errors.New("SelectPotion: invalid potion name")
	}
}
func (s *Service) CheckToken(token string) (*jwt.Token, error) {
	secret := os.Getenv("JWT_SECRET")
	issuer := os.Getenv("JWT_ISSUER")
	service := security.NewJWTService(secret, issuer)
	t, err := service.ValidateToken(token)
	if err != nil || !t.Valid {
		return &jwt.Token{}, authErr.InvalidToken("SRV-1010")
	}
	return t, nil
}
