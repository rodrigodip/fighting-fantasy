package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
)

func (ur *MongoUserRepo) Update(userID, field, newData string) error {
	filter := bson.M{"userId": userID}
	update := bson.M{"$set": bson.M{field: newData}}
	_, err := ur.coll.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}
	return nil
}
