package mongodb

import (
	"context"
	"fmt"
)

func (ur *mongoRepo) RegisterUser(id, name, email string, password []byte, role, status string) error {
	newUser := UserMongoEntity{
		UserID:   id,
		Name:     name,
		Email:    email,
		Password: password,
		Role:     role,
		Status:   status,
	}
	result, err := ur.coll.InsertOne(context.TODO(), newUser)
	if err != nil {
		return fmt.Errorf("Repository: %v", err)
	}
	// Prints the ID of the inserted document
	fmt.Printf("Document inserted with ID: %s\n", result.InsertedID)
	return nil
}
