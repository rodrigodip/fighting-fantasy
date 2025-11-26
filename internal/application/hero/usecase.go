package heroapp

import (
	"github.com/rodrigodip/fighting-fantasy/internal/domain/hero"
)

type HeroUseCase struct {
	service *hero.Service
	repo    Repository
}

func NewHeroUseCase(s *hero.Service, r Repository) *HeroUseCase {
	return &HeroUseCase{service: s, repo: r}
}

// func (uc *HeroUseCase) CreateHero(userID, heroName, potion string) (*hero.Hero, error) {
// 	// t, err := uc.HeroService.CheckToken(token)
// 	// if err != nil {
// 	// 	return &hero.Hero{}, fmt.Errorf("CreateHero: %v", err)
// 	// }
// 	// claims := t.Claims.(jwt.MapClaims)
// 	// userID := claims["user_id"].(string)
// 	if err := uc.service.ValidateInput(
// 		heroName, potion,
// 	); err != nil {
// 		return &hero.Hero{}, fmt.Errorf("CreateHero: %v", err)
// 	}
// 	newHero := hero.Hero{
// 		UserID:   userID,
// 		HeroName: heroName,
// 	}
// 	newHero, err := uc.service.SelectPotion(newHero, potion)
// 	if err != nil {
// 		return &hero.Hero{}, fmt.Errorf("CreateHero: %v", err)
// 	}
// 	if err := uc.repo.RegisterHero(newHero); err != nil {
// 		return &hero.Hero{}, fmt.Errorf("CreateHero: %v", err)
// 	}
// 	return &newHero, nil
// }
