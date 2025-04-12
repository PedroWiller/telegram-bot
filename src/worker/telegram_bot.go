package worker

import (
	"telegram-bot/src/config/db"
	"telegram-bot/src/repository"
	service "telegram-bot/src/service/telegram"
)

func Start() {
	dbMongo := db.ConnectMongoDB()
	userRepo := repository.NewUserRepository(dbMongo)
	telegramService := service.NewTelegramService(userRepo)
	go telegramService.SendMessageToUser()
	telegramService.Send()
}
