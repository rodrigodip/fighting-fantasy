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
func (s *Service) CreateHero(userID, heroName string) (*Hero, error) {
	if userID == "" {
		return &Hero{}, errors.New("CreateHero: invalid userID")
	}
	if heroName == "" {
		return &Hero{}, errors.New("CreateHero: heroName Required")
	}
	if len(heroName) < 3 {
		return &Hero{}, errors.New("CreateHero: heroName must have more than 3 digits")
	}
	newHero := &Hero{
		UserID:   userID,
		HeroName: heroName,
	}
	if err := s.service.RegisterHero(newHero.UserID, newHero.HeroName); err != nil {
		return &Hero{}, fmt.Errorf("CreateHero: %v", err)
	}
	return newHero, nil
}
