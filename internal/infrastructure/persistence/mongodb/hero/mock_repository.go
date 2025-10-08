package mongodb

import (
	"errors"

	"github.com/rodrigodip/fighting-fantasy/internal/domain/hero"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MockHeroRepository struct {
	heroes []*HeroMongoEntity
}

func NewMockHeroRepository() *MockHeroRepository {
	return &MockHeroRepository{
		heroes: make([]*HeroMongoEntity, 10),
	}
}

func (m *MockHeroRepository) RegisterHero(hero *hero.Hero) error {
	newHero := HeroMongoEntity{
		UserID:   hero.UserID,
		HeroName: hero.HeroName,
		Potions: Potions{
			Dexterity: hero.Potions.Dexterity,
			Strength:  hero.Potions.Strength,
			Fortune:   hero.Potions.Fortune,
		},
	}
	if newHero.ID == primitive.NilObjectID {
		newHero.ID = primitive.NewObjectID()
	}
	m.heroes = append(m.heroes, &newHero)
	return nil
}
func (m *MockHeroRepository) FindByOwner(userID string) (*hero.Hero, error) {
	var hero hero.Hero
	for _, storedHero := range m.heroes {
		if storedHero.UserID == userID {
			hero.UserID = storedHero.UserID
			hero.HeroName = storedHero.HeroName
			if storedHero.Potions.Dexterity {
				hero.Potions.Dexterity = true
			}
			if storedHero.Potions.Strength {
				hero.Potions.Strength = true
			}
			if storedHero.Potions.Fortune {
				hero.Potions.Fortune = true
			}
			return &hero, nil
		}
	}
	return nil, errors.New("MockUserRepo: Hero not found")
}
