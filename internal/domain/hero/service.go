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
func (s *Service) ValidateHero(h Hero) error {
	if h.UserID == "" {
		return errors.New("ValidadeHero: invalid userID")
	}
	if h.HeroName == "" {
		return errors.New("ValidadeHero: heroName Required")
	}
	if len(h.HeroName) < 3 {
		return errors.New("ValidadeHero: heroName must have more than 3 digits")
	}
	if (h.Potions.Dexterity || h.Potions.Fortune || h.Potions.Strength) == false {
		return errors.New("ValidadeHero: bad potion name: 'dexterity' || 'strength' || 'fortune' ")
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

// Those are types to build the hero
type (
	HeroBuilder interface {
		SetName(name string) HeroBuilder
		SetID(id string) HeroBuilder
		SelectPotion(potion string) HeroBuilder
		Build() Hero
	}
	ConcreteHeroBuilder struct {
		hero Hero
	}
)

func (b *ConcreteHeroBuilder) SetName(name string) HeroBuilder {
	b.hero.HeroName = name
	return b
}
func (b *ConcreteHeroBuilder) SetID(id string) HeroBuilder {
	b.hero.UserID = id
	return b
}
func (b *ConcreteHeroBuilder) SelectPotion(potion string) HeroBuilder {
	switch potion {
	case "dexterity":
		b.hero.Potions.Dexterity = true
		return b
	case "strength":
		b.hero.Potions.Strength = true
		return b
	case "fortune":
		b.hero.Potions.Fortune = true
		return b
	default:
		return b
	}
}
func (b *ConcreteHeroBuilder) Build() Hero {
	return b.hero
}
