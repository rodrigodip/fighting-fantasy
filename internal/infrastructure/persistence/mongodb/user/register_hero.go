package mongodb

import (
	"context"
	"fmt"

	"github.com/rodrigodip/fighting-fantasy/internal/domain/hero"
	heroEntity "github.com/rodrigodip/fighting-fantasy/internal/infrastructure/persistence/mongodb/hero"
)

func (hr *MongoUserRepo) RegisterHero(hero hero.Hero) error {
	newHero := heroEntity.HeroMongoEntity{
		UserID:      hero.UserID,
		HeroName:    hero.HeroName,
		CurrentLore: hero.CurrentLore,
		Stats: heroEntity.Stats{
			InitialDex:  hero.Stats.InitialDex,
			InitialHP:   hero.Stats.InitialHP,
			InitialLuck: hero.Stats.InitialLuck,
		},
		Inventory: heroEntity.Inventory{
			Equipment: hero.Inventory.Equipment,
			Gold:      hero.Inventory.Gold,
		},
		Potions: heroEntity.Potions{
			Dexterity: hero.Potions.Dexterity,
			Strength:  hero.Potions.Strength,
			Fortune:   hero.Potions.Fortune,
		},
	}
	result, err := hr.coll.InsertOne(context.TODO(), newHero)
	if err != nil {
		return fmt.Errorf("Repository: %v", err)
	}
	fmt.Printf("Hero created with ID:[%s] to UserID:[%s] \n", result.InsertedID, newHero.UserID)
	return nil
}
