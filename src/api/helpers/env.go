package envHelper

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetEnvValue(key string, defaultValue string) string {
	err := godotenv.Load(".env")

	if err != nil {
		log.Printf("Warning: Could not load .env file: %v", err)
	}

	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultValue
}
