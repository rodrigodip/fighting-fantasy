package mongodb

import "go.mongodb.org/mongo-driver/v2/mongo"

type mongoRepo struct {
	coll *mongo.Collection
}

func NewMongoRepository(database *mongo.Collection) *mongoRepo {
	return &mongoRepo{coll: database}
}
