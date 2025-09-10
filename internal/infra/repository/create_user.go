package repository

import (
	"context"
	"fmt"
)

func (ur *userRepository) CreateUser(name, email, pass string, age int) error {
	newUser := UserEntity{
		Name:     name,
		Email:    email,
		Password: pass,
		Age:      age,
	}
	//	defer func() {
	//		if err := client.Disconnect(context.TODO()); err != nil {
	//			panic(err)
	//		}
	//	}()
	coll := ur.dbConnection.Collection("user")
	result, err := coll.InsertOne(context.TODO(), newUser)
	if err != nil {
		panic(err)
	}
	// Prints the ID of the inserted document
	fmt.Printf("Document inserted with ID: %s\n", result.InsertedID)
	return nil
}
