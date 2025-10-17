package heroapp

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
	"github.com/rodrigodip/fighting-fantasy/internal/domain/hero"
)

type HeroUseCase struct {
	HeroService *hero.Service
}

func NewHeroUseCase(service *hero.Service) *HeroUseCase {
	return &HeroUseCase{HeroService: service}
}

func (uc *HeroUseCase) CreateHero(token, heroName, potion string) (*hero.Hero, error) {
	t, err := uc.HeroService.CheckToken(token)
	if err != nil {
		return &hero.Hero{}, fmt.Errorf("CreateHero: %v", err)
	}
	claims := t.Claims.(jwt.MapClaims)
	userID := claims["user_id"].(string)
	if err := uc.HeroService.ValidateInput(
		heroName, potion,
	); err != nil {
		return &hero.Hero{}, fmt.Errorf("CreateHero: %v", err)
	}
	newHero := hero.Hero{
		UserID:   userID,
		HeroName: heroName,
	}
	newHero, err = uc.HeroService.SelectPotion(newHero, potion)
	if err != nil {
		return &hero.Hero{}, fmt.Errorf("CreateHero: %v", err)
	}
	if err := uc.HeroService.Save(newHero); err != nil {
		return &hero.Hero{}, fmt.Errorf("CreateHero: %v", err)
	}
	return &newHero, nil
}
