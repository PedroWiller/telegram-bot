package worker

import (
	service "telegram-bot/src/service/telegram"
)

func Start() {
	telegramService := service.NewTelegramService()
	telegramService.Send()
}
