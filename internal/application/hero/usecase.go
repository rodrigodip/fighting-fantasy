package heroapp

import (
	"fmt"

	"github.com/rodrigodip/fighting-fantasy/internal/domain/hero"
)

type Repository interface {
	FindByOwner(userID string) (*hero.Hero, error)
	RegisterHero(hero hero.Hero) error
}
type HeroUseCase struct {
	service *hero.Service
	repo    Repository
}

func NewHeroUseCase(s *hero.Service, r Repository) *HeroUseCase {
	return &HeroUseCase{service: s, repo: r}
}

func (uc *HeroUseCase) CreateHero(userID, heroName, potion string) (*hero.Hero, error) {
	foundHero, err := uc.repo.FindByOwner(userID)
	if err == nil {
		if err := uc.service.HasHero(userID, *foundHero); err != nil {
			return &hero.Hero{}, fmt.Errorf("CreateHero: %v", err)
		}
	}
	if err := uc.service.ValidateInput(
		heroName, potion,
	); err != nil {
		return &hero.Hero{}, fmt.Errorf("CreateHero: %v", err)
	}
	newHero := hero.Hero{
		UserID:      userID,
		HeroName:    heroName,
		CurrentLore: 1,
		Stats: hero.Stats{
			InitialDex:  10,
			CurrentDex:  7,
			InitialHP:   10,
			CurrentHP:   6,
			InitialLuck: 10,
			CurrentLuck: 5,
		},
		Inventory: hero.Inventory{
			Equipment: []string{
				"Espada",
				"Armadura de Couro",
				"Mochila",
			},
			Backpack: &hero.Backpack{
				Provisions: 10,
				Gold:       30,
				Jewels:     0,
				Itens: []string{
					"Mapa para casa de Yaztronomo",
				},
			},
		},
	}
	newHero, err = uc.service.SelectPotion(newHero, potion)
	if err != nil {
		return &hero.Hero{}, fmt.Errorf("CreateHero: %v", err)
	}
	if err := uc.repo.RegisterHero(newHero); err != nil {
		return &hero.Hero{}, fmt.Errorf("CreateHero: %v", err)
	}
	return &newHero, nil
}
func (uc *HeroUseCase) FindByUser(userID string) (*hero.Hero, error) {
	foundHero, err := uc.repo.FindByOwner(userID)
	if err != nil {
		return &hero.Hero{}, fmt.Errorf("FindHeroByUser: %v", err)
	}
	return foundHero, nil
}
