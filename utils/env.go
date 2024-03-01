package utils

import (
	"os"

	"github.com/joho/godotenv"
)

func GetEnvKey(key string) string {
	godotenv.Load()
	return os.Getenv(key)
}
