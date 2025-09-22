package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
)

func (ur *userMongoRepo) SetVerified(userID string) error {
	filter := bson.M{"userId": userID}
	update := bson.M{"$set": bson.M{"status": "VERIFIED"}}
	_, err := ur.coll.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}
	return nil
}
