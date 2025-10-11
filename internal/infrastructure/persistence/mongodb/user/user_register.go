package mongodb

import (
	"context"
	"fmt"

	"github.com/rodrigodip/fighting-fantasy/internal/domain/user"
)

// RegisterUser persist a User on DB
func (ur *MongoUserRepo) RegisterUser(u user.User) error {
	newUser := UserMongoEntity{
		UserID:   u.UserID,
		Name:     u.Name,
		Email:    u.Email,
		Password: u.Password,
		Role:     u.Role,
		Status:   Status(u.Status),
	}
	result, err := ur.coll.InsertOne(context.TODO(), newUser)
	if err != nil {
		return fmt.Errorf("Repository: %v", err)
	}
	// Prints the ID of the inserted document
	fmt.Printf("Document inserted with ID: %s\n", result.InsertedID)
	return nil
}
