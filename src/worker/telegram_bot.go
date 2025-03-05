package worker

import (
	serviceGemini "telegram-bot/src/service/gemini"
	service "telegram-bot/src/service/telegram"
)

func Start() {
	geminiService := serviceGemini.NewGeminiService()
	telegramService := service.NewTelegramService(geminiService)
	telegramService.Send()
}
