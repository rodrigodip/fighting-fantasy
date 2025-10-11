package mongodb

import (
	"context"
	"fmt"
	"github.com/rodrigodip/fighting-fantasy/internal/domain/user"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func (ur *MongoUserRepo) FindById(userId string) (*usr.User, error) {
	var foundUser usr.User
	err := ur.coll.FindOne(context.TODO(), bson.M{"userId": userId}).Decode(&foundUser)
	if err != nil {
		return &usr.User{}, fmt.Errorf("Repository: %v", err)
	}
	return &foundUser, nil
}
