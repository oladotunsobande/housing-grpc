package config

import (
	"log"
	"os"
	"path"

	"github.com/joho/godotenv"
)

// Secrets Struct for all environment variables
type Secrets struct {
	Environment     string
	ApplicationPort string
	ApplicationName string
}

func init() {
	if err := godotenv.Load(path.Join(".env")); err != nil {
		log.Fatal("Error loading .env file")
	}
	log.Println(".env loaded")
}

// GetSecrets Get all loaded secrets
func GetSecrets() Secrets {
	var secrets Secrets

	secrets.Environment = os.Getenv("GO_ENV")
	secrets.ApplicationPort = os.Getenv("APP_PORT")
	secrets.ApplicationName = os.Getenv("APP_NAME")

	return secrets
}
