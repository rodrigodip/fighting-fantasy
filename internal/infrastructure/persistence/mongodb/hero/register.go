package mongodb

import (
	//	"context"
	//	"fmt"

	"github.com/rodrigodip/fighting-fantasy/internal/domain/hero"
)

//	func (hr *MongoHeroRepo) RegisterHero(hero hero.Hero) error {
//		newHero := HeroMongoEntity{
//			UserID:      hero.UserID,
//			HeroName:    hero.HeroName,
//			CurrentLore: hero.CurrentLore,
//			Stats: Stats{
//				InitialDex:  hero.Stats.InitialDex,
//				InitialHP:   hero.Stats.InitialHP,
//				InitialLuck: hero.Stats.InitialLuck,
//			},
//			Inventory: Inventory{
//				Equipment: hero.Inventory.Equipment,
//				Gold:      hero.Inventory.Gold,
//			},
//			Potions: Potions{
//				Dexterity: hero.Potions.Dexterity,
//				Strength:  hero.Potions.Strength,
//				Fortune:   hero.Potions.Fortune,
//			},
//		}
//		result, err := hr.coll.InsertOne(context.TODO(), newHero)
//		if err != nil {
//			return fmt.Errorf("Repository: %v", err)
//		}
//		fmt.Printf("Hero created with ID:[%s] to UserID:[%s] \n", result.InsertedID, newHero.UserID)
//		return nil
//	}
func (hr *MongoHeroRepo) FindByOwner(userID string) (*hero.Hero, error) {
	return &hero.Hero{}, nil
}
