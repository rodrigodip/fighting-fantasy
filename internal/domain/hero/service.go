package hero

import (
	"errors"
	"fmt"
)

type Service struct {
	service Repository
}

func NewHeroService(r Repository) *Service {
	return &Service{service: r}
}

// ValidateHero enforces hero business rules
func (s *Service) ValidateInput(userID, heroName, potion string) error {
	if userID == "" {
		return errors.New("ValidadeHero: userID is Required")
	}
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
		return errors.New("ValidadeHero: potion is Required")
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

// SelectPotion attach potion to Hero
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
