package repository

import (
	"context"
	"fmt"
)

func (ur *userRepository) CreateUser(id, name, email, pass string, role []string) error {
	newUser := UserEntity{
		UserID:   id,
		Name:     name,
		Email:    email,
		Password: pass,
		Roles:    role,
	}

	coll := ur.dbConnection.Collection("user")
	result, err := coll.InsertOne(context.TODO(), newUser)
	if err != nil {
		panic(err)
	}
	// Prints the ID of the inserted document
	fmt.Printf("Document inserted with ID: %s\n", result.InsertedID)
	return nil
}
