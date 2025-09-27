package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

// CreateUserIndexes creates constraints(index) on User Collection
func CreateUserIndexes(ctx context.Context, dbName string, client *mongo.Client) error {
	userColl := client.Database(dbName).Collection("user")
	indexModel := mongo.IndexModel{
		Keys:    bson.D{{Key: "email", Value: 1}},
		Options: options.Index().SetUnique(true),
	}
	_, err := userColl.Indexes().CreateOne(ctx, indexModel)
	if err != nil {
		return err
	}
	return nil
}
