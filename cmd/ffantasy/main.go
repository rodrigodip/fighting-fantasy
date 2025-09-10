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
	database := repository.NewUserRepository(db)
	service := usecase.NewUserRopository(database)

	output, err := service.CreateUser("roberta", "roberta@gmail.com", "Pipa@121982", 42)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(output)
}
