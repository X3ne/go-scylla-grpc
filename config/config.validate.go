package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func ValidateEnvs() map[string]string {
	err:= godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	envs := map[string]string{
		"HOST": os.Getenv("HOST"),
		"PORT": os.Getenv("PORT"),
		"DB_HOSTS": os.Getenv("DB_HOSTS"),
		"DB_KEYSPACE": os.Getenv("DB_KEYSPACE"),
		"JWT_SECRET_KEY": os.Getenv("JWT_SECRET_KEY"),
		"JWT_TOKEN_DURATION": os.Getenv("JWT_TOKEN_DURATION"),
	}

	for key, value := range envs {
		if value == "" {
			log.Fatalf("Missing %s environment variable", key)
		}
	}

	return envs
}
