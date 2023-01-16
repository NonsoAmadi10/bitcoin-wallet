package utils

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

func GetEnv(key string) string {

	return os.Getenv(key)
}

type Response struct {
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}
