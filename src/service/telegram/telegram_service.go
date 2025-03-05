package service

import (
	"fmt"
	"log"

	"telegram-bot/src/config/env"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type TelegramService struct{}

func NewTelegramService() *TelegramService {
	return &TelegramService{}
}

func (s *TelegramService) Send() {
	bot, err := tgbotapi.NewBotAPI(env.TelegramApiToken)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	fmt.Println("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	var responseMsg string
	for update := range updates {
		if update.Message != nil {
			responseMsg = "Ola"

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, responseMsg)
			msg.ReplyToMessageID = update.Message.MessageID

			if _, err := bot.Send(msg); err != nil {
				log.Fatalf(err.Error())
			}
		}
	}
}
