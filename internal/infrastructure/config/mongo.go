package config

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func NewMongoDBConnection(ctx context.Context) (*mongo.Database, error) {
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("Set your 'MONGODB_URI' environment variable.")
	}

	client, err := mongo.Connect(options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	//TODO: REFACTOR: This block creates User Collection's constraints (index)
	userColl := client.Database("ffantasy-db").Collection("user")
	indexModel := mongo.IndexModel{
		Keys:    bson.D{{Key: "email", Value: 1}},
		Options: options.Index().SetUnique(true),
	}
	_, err = userColl.Indexes().CreateOne(ctx, indexModel)
	if err != nil {
		return nil, err
	}

	//TODO: REFACTOR
	if err := client.Ping(ctx, nil); err != nil {
		return nil, err
	}
	//NOTE:"ffantasy-db" must be a SECRETE
	return client.Database("ffantasy-db"), nil
}
