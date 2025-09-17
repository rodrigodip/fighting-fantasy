package config

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type Config struct {
	MongoURI  string
	DBName    string
	JWTSecret string
	JWTIssuer string
	HTTPPort  string
}

func LoadConfig() *Config {
	return &Config{
		MongoURI:  os.Getenv("MONGO_URI"),
		DBName:    os.Getenv("MONGO_DB"),
		JWTSecret: os.Getenv("JWT_SECRET"),
		JWTIssuer: os.Getenv("JWT_ISSUER"),
		HTTPPort:  getEnv("HTTP_PORT", "8080"),
	}
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
func NewMongoDBConnection(ctx context.Context) (*mongo.Database, error) {
	uri := LoadConfig().MongoURI
	dbName := LoadConfig().DBName
	if uri == "" {
		log.Fatal("Set your 'MONGODB_URI' environment variable.")
	}

	client, err := mongo.Connect(options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	//TODO: REFACTOR: This block creates User Collection's constraints (index)
	userColl := client.Database(dbName).Collection("user")
	indexModel := mongo.IndexModel{
		Keys:    bson.D{{Key: "email", Value: 1}},
		Options: options.Index().SetUnique(true),
	}
	_, err = userColl.Indexes().CreateOne(ctx, indexModel)
	if err != nil {
		return nil, err
	}

	if err := client.Ping(ctx, nil); err != nil {
		return nil, err
	}
	return client.Database(dbName), nil
}
