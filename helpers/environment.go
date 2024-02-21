package helpers

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetEnvParam(parameter string, defaultVal string) string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("No env parametes set")
		os.Exit(1)
	}

	param := os.Getenv(parameter)
	if param != "" {
		return param
	}

	return defaultVal
}
