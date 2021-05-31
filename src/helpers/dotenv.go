package helpers

import (
	"github.com/joho/godotenv"
)

func LoadDotEnv() {
	if godotenv.Load() != nil {
		panic("Error loading .env file")
	}
}
