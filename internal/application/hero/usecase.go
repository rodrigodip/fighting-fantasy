package heroapp

import "github.com/rodrigodip/fighting-fantasy/internal/domain/hero"

type HeroUseCase struct {
	HeroService *hero.Service
}

func NewHeroUseCase(service *hero.Service) *HeroUseCase {
	return &HeroUseCase{HeroService: service}
}

func (uc *HeroUseCase) CreateHero(userID, heroName string) (*hero.Hero, error) {
	return uc.HeroService.CreateHero(userID, heroName)
}
