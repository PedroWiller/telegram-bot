package service

import (
	"log"

	"telegram-bot/src/config/client"
	"telegram-bot/src/config/env"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"google.golang.org/genai"
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

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {

			sendedMessage := update.Message.Text

			model := "gemini-1.5-pro-002" // Ou "gemini-1.5-pro-002" para melhor qualidade
			result, err := client.GlobalClient.Models.GenerateContent(client.Ctx, model, genai.Text(sendedMessage), nil)
			response := "Desculpe, não consegui entender o que você disse. Poderia repetir?"
			if err == nil {
				response, _ = result.Text()
			}

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, response)
			msg.ReplyToMessageID = update.Message.MessageID

			if _, err := bot.Send(msg); err != nil {
				log.Fatalf(err.Error())
			}
		}
	}
}
