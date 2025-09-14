package mongodb

import (
	"context"
	"fmt"
)

func (ur *userMongoRepo) RegisterUser(id, name, email string, password []byte, role []string) error {
	newUser := UserMongoEntity{
		UserID:   id,
		Name:     name,
		Email:    email,
		Password: password,
		Roles:    role,
	}
	result, err := ur.coll.InsertOne(context.TODO(), newUser)
	if err != nil {
		return fmt.Errorf("repository.CreateUser(): %v", err)
	}
	// Prints the ID of the inserted document
	fmt.Printf("Document inserted with ID: %s\n", result.InsertedID)
	return nil
}
