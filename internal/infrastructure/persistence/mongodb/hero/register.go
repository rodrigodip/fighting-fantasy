package mongodb

import (
	"context"
	"fmt"

	"github.com/rodrigodip/fighting-fantasy/internal/domain/hero"
)

func (hr *MongoHeroRepo) RegisterHero(hero hero.Hero) error {
	newHero := HeroMongoEntity{
		UserID:      hero.UserID,
		HeroName:    hero.HeroName,
		CurrentLore: hero.CurrentLore,
		Stats: Stats{
			InitialDex:  hero.Stats.InitialDex,
			InitialHP:   hero.Stats.InitialHP,
			InitialLuck: hero.Stats.InitialLuck,
			CurrentDex:  hero.Stats.CurrentDex,
			CurrentHP:   hero.Stats.CurrentHP,
			CurrentLuck: hero.Stats.CurrentLuck,
		},
		Inventory: Inventory{
			Equipment: hero.Inventory.Equipment,
			Backpack: Backpack{
				Provisions: hero.Inventory.Backpack.Provisions,
				Gold:       hero.Inventory.Backpack.Gold,
				Jewels:     hero.Inventory.Backpack.Jewels,
				Itens:      hero.Inventory.Backpack.Itens,
			},
		},
		Potions: Potions{
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
