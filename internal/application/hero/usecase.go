package heroapp

import (
	"fmt"

	"github.com/rodrigodip/fighting-fantasy/internal/domain/hero"
)

type HeroUseCase struct {
	HeroService *hero.Service
}

func NewHeroUseCase(service *hero.Service) *HeroUseCase {
	return &HeroUseCase{HeroService: service}
}

func (uc *HeroUseCase) CreateHero(userID, heroName, potion string) (*hero.Hero, error) {

	builder := &hero.ConcreteHeroBuilder{}
	newHero := builder.SetID(userID).SetName(heroName).SelectPotion(potion).Build()
	if err := uc.HeroService.ValidateHero(newHero); err != nil {
		return &hero.Hero{}, fmt.Errorf("CreateHero: %v", err)
	}
	if err := uc.HeroService.Save(newHero); err != nil {
		return &hero.Hero{}, fmt.Errorf("CreateHero: %v", err)
	}
	return &newHero, nil
}
