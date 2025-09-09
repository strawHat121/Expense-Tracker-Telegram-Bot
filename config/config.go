package config

import "os"

func GetAPIKey() string {
	return os.Getenv("TELEGRAM_BOT_API_KEY")
}
