package config

import (
	"github.com/joho/godotenv"
)

func LoadEnvironmentVariables(filename string) {
	godotenv.Load(filename)
}
