package env

import (
	"fmt"
	"os"
)

var (
	PORT             string
	TelegramApiToken string
	GeminiApiToken   string
	MongoUri         string
)

func Start() error {
	PORT = os.Getenv("PORT")
	if PORT == "" {
		PORT = "5001"
	}

	TelegramApiToken = os.Getenv("TELEGRAM_API_TOKEN")
	if TelegramApiToken == "" {
		return fmt.Errorf("Error to load TELEGRAM_API_TOKEN")
	}
	GeminiApiToken = os.Getenv("GEMINI_API_KEY")
	if GeminiApiToken == "" {
		return fmt.Errorf("Error to load GEMINI_API_KEY")
	}

	MongoUri = os.Getenv("MONGO_URI")
	if MongoUri == "" {
		MongoUri = "mongodb://159.89.36.80:27017/telegram"
	}

	return nil
}
