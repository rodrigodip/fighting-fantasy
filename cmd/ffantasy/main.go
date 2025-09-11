package main

import (
	"context"
	"fmt"
	"log"

	"github.com/rodrigodip/fighting-fantasy/internal/infra/data_base/mongoDB"
	"github.com/rodrigodip/fighting-fantasy/internal/infra/repository"
	"github.com/rodrigodip/fighting-fantasy/internal/use_case/user"
)

func main() {
	db, err := mongodb.NewMongoDBConnection(context.Background())
	if err != nil {
		log.Fatalf(
			"Error trying to connect to database, error=%s \n",
			err.Error())
		return
	}
	defer func() {
		if err := db.Client().Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	database := repository.NewUserRepository(db)
	service := usecase.NewUserRepository(database)

	output, err := service.CreateUser("roberta", "roberta@gmail.com", "Pipa@121982")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(output)
}
