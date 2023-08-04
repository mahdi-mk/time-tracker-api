package env

import "os"

// Get retrieves the value of the environment variable named by the key.
// If the variable does not exist, it returns the provided fallback value.
func Get(key string, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return fallback
}
