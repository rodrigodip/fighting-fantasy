package mongodb

import (
	"go.mongodb.org/mongo-driver/v2/mongo"
)

// type UserMongoRepository interface {
// 	CreateUser(id, name, email, pass string, role []string) error
// 	//GetUser(id string)(User, error)
// 	//GetUserByEmailAndPass(email, password string) (User, error)
// 	//UpdateUser(id, name, email string, age int)(User, error)
// 	//DeleteUser(id string) error
// }

type userMongoRepo struct {
	coll *mongo.Collection
}

func NewUserRepository(database *mongo.Collection) *userMongoRepo {
	return &userMongoRepo{coll: database}
}
