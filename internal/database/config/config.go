package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	// StrinConnect is the connection string to MySQL
	StrinConnect = ""

	// APIAddr is the API port address
	ApiAddr = 0
)

func LoadEnv() {
	var err error

	if err = godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	ApiAddr, err = strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		ApiAddr = 9000
	}

	StrinConnect = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USUARIO"),
		os.Getenv("BD_PW"),
		os.Getenv("DB_NAME"),
	)
}
