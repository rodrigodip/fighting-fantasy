package mongodb

import (
	"context"
	"fmt"
	"log"

	"github.com/rodrigodip/fighting-fantasy/internal/domain/hero"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func (hr *MongoHeroRepo) FindByOwner(userID string) (*hero.Hero, error) {
	var foundHero hero.Hero
	err := hr.coll.FindOne(context.TODO(), bson.M{"userId": userID}).Decode(&foundHero)
	if err != nil {
		return &hero.Hero{}, fmt.Errorf("repository.FindByOwner(): %v", err)
	}
	log.Printf("repository.FindByOwner(): %s", foundHero.HeroName)
	return &foundHero, nil
}
