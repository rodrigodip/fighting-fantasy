package hero

import (
	"errors"
)

type Service struct {
	service Repository
}

func NewHeroService() *Service {
	return &Service{}
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
