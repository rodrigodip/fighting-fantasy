package mongodb

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type MongoAdventureRepo struct {
	coll *mongo.Collection
}

func NewMongoAdventureRepository(database *mongo.Collection) *MongoAdventureRepo {
	return &MongoAdventureRepo{coll: database}
}

type Segment struct {
	ID            primitive.ObjectID `bson:"_id"`
	LoreID        int                `bson:"lore_id"`
	Text          string             `bson:"text"`
	Combat        *Combat            `bson:"combat,omitempty"`
	LuckTest      *LuckTest          `bson:"luck_test,omitempty"`
	AttributeTest *AttributeTest     `bson:"attribute_test,omitempty"`
	ItemCheck     *ItemCheck         `bson:"item_check,omitempty"`
	Effects       []Effect           `bson:"effects"`
	Choices       []Choice           `bson:"choices"`
	IsTerminal    bool               `bson:"is_terminal"`
}

type Combat struct {
	Escape      *Escape  `bson:"escape,omitempty"`
	Penalty     *Penalty `bson:"penalty,omitempty"`
	Waves       []Wave   `bson:"waves"`
	VictoryGoto int      `bson:"victory_goto"`
	DefeatGoto  *int     `bson:"defeat_goto,omitempty"`
}

type Escape struct {
	Available   bool `bson:"available"`
	AfterRounds int  `bson:"after_rounds"`
	Goto        int  `bson:"goto"`
	StaminaCost int  `bson:"stamina_cost"`
}

type Penalty struct {
	HeroAttackModifier int    `bson:"hero_attack_modifier"`
	Description        string `bson:"description"`
}

type Wave struct {
	Description string  `bson:"description"`
	Mode        string  `bson:"mode"`
	Enemies     []Enemy `bson:"enemies"`
}

type Enemy struct {
	Name    string `bson:"name"`
	Skill   int    `bson:"skill"`
	Stamina int    `bson:"stamina"`
}

type LuckTest struct {
	Description string  `bson:"description"`
	Success     Outcome `bson:"success"`
	Failure     Outcome `bson:"failure"`
}

type Outcome struct {
	Goto      int        `bson:"goto"`
	Effects   []Effect   `bson:"effects"`
	Condition *Condition `bson:"condition,omitempty"`
}

type Condition struct {
	MustSurvive bool `bson:"must_survive"`
}

type AttributeTest struct {
	// Extensível futuramente
}

type ItemCheck struct {
	Item        string `bson:"item"`
	SuccessGoto int    `bson:"success_goto"`
	FailureGoto int    `bson:"failure_goto"`
}

type Effect struct {
	Type      string `bson:"type"`
	Operation string `bson:"operation"`
	Value     int    `bson:"value,omitempty"`
}

type Choice struct {
	Description *string `bson:"description,omitempty"`
	Goto        int     `bson:"goto"`
}
