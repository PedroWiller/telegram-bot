package env

import (
	"fmt"
	"os"
)

var (
	PORT             string
	TelegramApiToken string
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

	return nil
}
