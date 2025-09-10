package repository

import "go.mongodb.org/mongo-driver/v2/mongo"

// const (
//
//	MONGODB_USER_DB = "MONGODB_USER_DB"
//
// )
type UserEntity struct {
	Name     string `bson:"name,omitempty"`
	Email    string `bson:"email,omitempty"`
	Password string `bson:"password,omitempty"`
	Age      int    `bson:"age,omitempty"`
}

type UserRepository interface {
	CreateUser(name, email, pass string, age int) error
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
