package utils

import "github.com/joho/godotenv"

func LoadEnv() {
	err := godotenv.Load(".env")
	FailOnError(err, "Unable to load .env file")
}
