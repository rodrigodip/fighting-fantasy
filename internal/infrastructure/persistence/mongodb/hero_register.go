package mongodb

import (
	"context"
	"fmt"
)

func (hr *mongoRepo) RegisterHero(userID, heroName string) error {
	newHero := HeroMongoEntity{
		UserID:      userID,
		HeroName:    heroName,
		CurrentLore: 0,
		Stats: Stats{
			InitialDex:  6,
			InitialHP:   12,
			InitialLuck: 6,
		},
		Inventory: Inventory{
			Equipment: []string{"espada_simples", "armadura_couro"},
			Gold:      30,
		},
	}
	result, err := hr.coll.InsertOne(context.TODO(), newHero)
	if err != nil {
		return fmt.Errorf("Repository: %v", err)
	}
	fmt.Printf("Hero created with ID:[%s] to UserID:[%s] \n", result.InsertedID, userID)
	return nil
}
