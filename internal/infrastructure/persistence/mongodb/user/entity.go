package mongodb

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type UserMongoEntity struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	UserID    string             `bson:"userId"`
	Name      string             `bson:"name"`
	Email     string             `bson:"email"`
	Password  []byte
	Status    string    `bson:"status"`
	CreatedAt time.Time `bson:"createdAt"`
	Profile   Profile   `bson:"profile"`
	Role      string    `bson:"role"`
}
type Profile struct {
	Age    int    `bson:"age"`
	Gender string `bson:"gender"`
}

type userMongoRepo struct {
	coll *mongo.Collection
}

func NewUserRepository(database *mongo.Collection) *userMongoRepo {
	return &userMongoRepo{coll: database}
}
