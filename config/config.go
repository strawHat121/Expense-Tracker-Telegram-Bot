package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetAPIKey() string {

	err := godotenv.Load()

	if err != nil {
		log.Fatal(err)
	}

	return os.Getenv("TELEGRAM_BOT_API_KEY")
}
