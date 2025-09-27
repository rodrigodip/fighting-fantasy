package mongodb

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserMongoEntity struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	UserID    string             `bson:"userId"`
	Name      string             `bson:"name"`
	Email     string             `bson:"email"`
	Password  []byte
	Status    string              `bson:"status"`
	CreatedAt primitive.Timestamp `bson:"createdAt"`
	Profile   Profile             `bson:"profile"`
	Role      string              `bson:"role"`
}
type Profile struct {
	Age    int    `bson:"age"`
	Gender string `bson:"gender"`
}
