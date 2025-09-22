package mongodb

import (
	"context"
	"fmt"
	"github.com/rodrigodip/fighting-fantasy/internal/domain/user"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func (ur *userMongoRepo) FindById(userId string) (*user.User, error) {
	var foundUser user.User
	err := ur.coll.FindOne(context.TODO(), bson.M{"userId": userId}).Decode(&foundUser)
	if err != nil {
		return &user.User{}, fmt.Errorf("Repository: %v", err)
	}
	return &foundUser, nil
}

