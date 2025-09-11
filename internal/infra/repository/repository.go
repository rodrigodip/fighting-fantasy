package repository

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Profile struct {
	Age    int    `bson:"age"`
	Gender string `bson:"gender"`
}
type UserEntity struct {
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

type UserRepository interface {
	CreateUser(id, name, email, pass string, role []string) error
	//GetUser(id string)(User, error)
	//GetUserByEmailAndPass(email, password string) (User, error)
	//UpdateUser(id, name, email string, age int)(User, error)
	//DeleteUser(id string) error
}

type userRepository struct {
	dbConnection *mongo.Database
}

func NewUserRepository(database *mongo.Database) *userRepository {
	return &userRepository{dbConnection: database}
}
