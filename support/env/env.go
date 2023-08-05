package env

import (
	"os"

	"github.com/gofiber/fiber/v2/log"
	"github.com/joho/godotenv"
)

// Init initializes the "env" support.
func Init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("An error occurred while loading environment variables", err)
	}
}

// Get retrieves the value of the environment variable named by the key.
// If the variable does not exist, it returns the provided fallback value.
func Get(key string, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return fallback
}
