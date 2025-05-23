package service

import (
	"fmt"
	"log"
	"time"

	"telegram-bot/src/config/client"
	"telegram-bot/src/config/env"
	"telegram-bot/src/factory"
	"telegram-bot/src/repository/interfaces"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"google.golang.org/genai"
)

type TelegramService struct {
	Messages       []*genai.Part
	userRepository interfaces.UserRepository
}

func NewTelegramService(userRepository interfaces.UserRepository) *TelegramService {
	return &TelegramService{
		userRepository: userRepository,
	}
}

func (c *TelegramService) AddMessage(role, text string) {
	part := &genai.Part{Text: fmt.Sprintf("%s: %s", role, text)}
	c.Messages = append(c.Messages, part)

	maxMessages := 1
	if len(c.Messages) > maxMessages {
		c.Messages = c.Messages[1:]
	}
}

func StartBot() *tgbotapi.BotAPI {
	bot, err := tgbotapi.NewBotAPI(env.TelegramApiToken)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = false

	return bot
}

func (s *TelegramService) Send() {
	bot := StartBot()

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		var message string
		if update.Message != nil {
			userSended := update.Message
			sendedMessage := update.Message.Text

			chatID := update.Message.Chat.ID

			userModel := factory.CreateUser(userSended.From.ID, userSended.From.FirstName, userSended.From.UserName, chatID)
			returnedUser, _ := s.userRepository.GetBydId(userModel.ExternalId)
			if returnedUser.ExternalId == 0 {
				message = "Vc, ainda não possui assinatura, deseja receber notficias, 1 para aceitar"
				sendBotMessage(chatID, bot, message)
				_ = s.userRepository.Save(userModel)

				continue
			}

			if sendedMessage == "1" {
				userModel.SendedMessage = true
				_ = s.userRepository.Save(userModel)
				message = "Assinatura ativa, voçe receberá mensagens diárias"
				sendBotMessage(chatID, bot, message)

				s.PostMessages(chatID)

				continue
			}

		}
	}
}

func GeminiMessage(contents []*genai.Content) string {
	model := "gemini-1.5-pro-002" // Ou "gemini-1.5-pro-002" para melhor qualidade
	result, _ := client.GlobalClient.Models.GenerateContent(
		client.Ctx,
		model,
		contents,
		nil,
	)

	response, _ := result.Text()

	return response
}

func sendBotMessage(chatId int64, bot *tgbotapi.BotAPI, message string) {
	msg := tgbotapi.NewMessage(chatId, message)
	if _, err := bot.Send(msg); err != nil {
		log.Fatalf(err.Error())
	}
}

func (s *TelegramService) CreateNewMessage(message string) string {
	s.AddMessage("system", message)

	contents := make([]*genai.Content, 1)
	contents[0] = &genai.Content{
		Parts: s.Messages,
	}
	return GeminiMessage(contents)
}

func (s *TelegramService) SendMessageToUser() {
	for {
		users, err := s.userRepository.GetAll()
		if err != nil {
			fmt.Sprintln(err.Error())
			return
		}

		for _, user := range users {
			s.PostMessages(user.ChatId)
		}

		time.Sleep(time.Hour * 12)
	}
}

func (s *TelegramService) PostMessages(chatId int64) {
	messages := make([]string, 4)
	messages[0] = s.CreateNewMessage("Mensagem aleatoria de bom dia")
	messages[1] = s.CreateNewMessage("Por favor gerar um trecho simplificado do clean code")
	messages[2] = s.CreateNewMessage("Por favor, me informe uma curiosidade sobre BI")
	messages[3] = s.CreateNewMessage("Por favor, me informe uma curiosidade sobre Programação")

	sendBotMessage(chatId, StartBot(), messages[0])
	sendBotMessage(chatId, StartBot(), messages[1])
	sendBotMessage(chatId, StartBot(), messages[2])
	sendBotMessage(chatId, StartBot(), messages[3])
}
