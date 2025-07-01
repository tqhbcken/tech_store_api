package config

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnvVar() {
	err := godotenv.Load()
	if err != nil {
		// In a production/docker environment, it's normal for the .env file to be missing.
		// Environment variables will be loaded from the system.
		log.Println("Warning: .env file not found. Using system environment variables.")
	}
}
