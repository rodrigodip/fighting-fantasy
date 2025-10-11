package mongodb

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type MongoHeroRepo struct {
	coll *mongo.Collection
}

func NewMongoHeroRepository(database *mongo.Collection) *MongoHeroRepo {
	return &MongoHeroRepo{coll: database}
}

type HeroMongoEntity struct {
	ID              primitive.ObjectID `bson:"_id,omitempty"`
	UserID          string             `bson:"userId"`
	HeroName        string             `bson:"hero_name"`
	CurrentLore     int                `bson:"curnt_lore"`
	Stats           Stats              `bson:"stats"`
	Inventory       Inventory          `bson:"inventory"`
	Potions         Potions            `bson:"potions"`
	History         []History          `bson:"history"`
	MonsterDefeated []string           `bson:"Monster_Defeated"`
}
type Stats struct {
	InitialDex  int `bson:"init_dex"`
	CurrentDex  int `bson:"curnt_dex"`
	InitialHP   int `bson:"init_health"`
	CurrentHP   int `bson:"curnt_health"`
	InitialLuck int `bson:"init_luck"`
	CurrentLuck int `bson:"curnt_luck"`
}
type Inventory struct {
	Equipment  []string `bson:"equipment"`
	Backpack   []string `bson:"backpack"`
	Provisions int      `bson:"provisions"`
	Gold       int      `bson:"gold"`
	Jewels     int      `bson:"jewls"`
}
type Potions struct {
	Dexterity bool `bson:"dexterity"`
	Strength  bool `bson:"strength"`
	Fortune   bool `bson:"fortune"`
}
type History struct {
	Lore      int                 `bson:"lore"`
	Timestamp primitive.Timestamp `bson:"timestamp"`
	Choices   []string            `bson:"choices"`
}
