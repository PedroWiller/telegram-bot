package service

import (
	"telegram-bot/src/config/client"

	"google.golang.org/genai"
)

type GeminiService struct{}

func NewGeminiService() *GeminiService {
	return &GeminiService{}
}

func (s *GeminiService) NewMessage(prompt string, content string) (string, error) {
	model := "gemini-1.5-flash-002" // Ou "gemini-1.5-pro-002" para melhor qualidade
	result, err := client.GlobalClient.Models.GenerateContent(client.Ctx, model, genai.Text(prompt), nil)
	if err != nil {
		return "", err
	}

	// Exibir resposta
	return result.Text()
}
