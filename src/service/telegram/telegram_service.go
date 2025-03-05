package service

import (
	"fmt"
	"log"

	"telegram-bot/src/config/env"
	service "telegram-bot/src/service/gemini"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type TelegramService struct {
	geminiService *service.GeminiService
}

func NewTelegramService(geminiService *service.GeminiService) *TelegramService {
	return &TelegramService{
		geminiService: geminiService,
	}
}

func (s *TelegramService) Send() {
	bot, err := tgbotapi.NewBotAPI(env.TelegramApiToken)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {

			sendedMessage := update.Message.Text
			senderID := fmt.Sprintf("%d", update.Message.Chat.ID)
			responseMessage, err := s.geminiService.NewMessage(sendedMessage, senderID)
			if err != nil {
				responseMessage = err.Error()
			}

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, responseMessage)
			msg.ReplyToMessageID = update.Message.MessageID

			if _, err := bot.Send(msg); err != nil {
				log.Fatalf(err.Error())
			}
		}
	}
}
