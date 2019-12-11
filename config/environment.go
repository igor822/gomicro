package config

import "os"

var (
	// Port to be listened by application
	Port string
)

func getEnv(key, defaultValue string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return defaultValue
}

func init() {
	Port = getEnv("PORT", "8080")
}
