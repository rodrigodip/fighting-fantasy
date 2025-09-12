package mongodb

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserMongoEntity struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	UserID    string             `bson:"userId"`
	Name      string             `bson:"name"`
	Email     string             `bson:"email"`
	Password  string
	Status    string    `bson:"status"`
	CreatedAt time.Time `bson:"createdAt"`
	Profile   Profile   `bson:"profile"`
	Roles     []string  `bson:"roles"`
}
type Profile struct {
	Age    int    `bson:"age"`
	Gender string `bson:"gender"`
}
