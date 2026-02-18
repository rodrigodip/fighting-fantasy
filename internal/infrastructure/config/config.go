package config

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/sessions"
	userRepository "github.com/rodrigodip/fighting-fantasy/internal/infrastructure/persistence/mongodb/user"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type Config struct {
	MongoURI      string
	DBName        string
	JWTSecret     string
	JWTIssuer     string
	SessionSecret string
	SessionOpt    *sessions.Options
	HTTPPort      string

	SMTPHost     string
	SMTPPort     string
	SMTPUsername string
	SMTPPassword string
	SMTPFrom     string

	TemplatesPath string
}

func LoadConfig() *Config {
	return &Config{
		MongoURI:      os.Getenv("MONGO_URI"),
		DBName:        os.Getenv("MONGO_DB"),
		JWTSecret:     os.Getenv("JWT_SECRET"),
		JWTIssuer:     os.Getenv("JWT_ISSUER"),
		SessionSecret: os.Getenv("SESSION_SECRET"),
		SessionOpt:    setEnvironment(),
		HTTPPort:      getEnv("HTTP_PORT", "8080"),

		SMTPHost:      os.Getenv("SMTP_HOST"),
		SMTPPort:      os.Getenv("SMTP_PORT"),
		SMTPUsername:  os.Getenv("SMTP_USERNAME"),
		SMTPPassword:  os.Getenv("SMTP_PASSWORD"),
		SMTPFrom:      os.Getenv("SMTP_FROM"),
		TemplatesPath: os.Getenv("WEB_TEMPLATE_PATH"),
	}
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}

func setEnvironment() *sessions.Options {
	s := &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7, // 7 days
		HttpOnly: true,
		Secure:   true, // Set true in production with HTTPS
		SameSite: http.SameSiteLaxMode,
	}
	envType := os.Getenv("SET_ENVIRONMENT")
	if envType == "DEV" {
		s.Secure = false
	}
	log.Printf("=== YOU ARE IN %s ENVIRONMENT ===", envType)
	return s
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
	//Collection's constraints (index)
	if err := userRepository.CreateUserIndexes(ctx, dbName, client); err != nil {
		return nil, err
	}
	//Ping DB
	if err := client.Ping(ctx, nil); err != nil {
		return nil, err
	}
	log.Print("Visit http://localhost:8080/swagger/index.html for project SWAGGER")
	return client.Database(dbName), nil
}
