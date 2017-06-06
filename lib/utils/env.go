package utils

import "os"

// Get environment var with a fallback
func GetEnv(key string, fallback string) string {
	value := os.Getenv(key)

	if len(value) == 0 {
		return fallback
	}

	return value
}
