package service

import (
	"context"
	"errors"
	"os"

	"google.golang.org/genai"
)

type GeminiService struct{}

func NewGeminiService() *GeminiService {
	return &GeminiService{}
}

func (s *GeminiService) NewMessage(prompt string, content string) (string, error) {
	apiKey := os.Getenv("GEMINI_API_KEY")
	if apiKey == "" {
		return "", errors.New("GEMINI_API_KEY não definida")
	}

	// Criar cliente Gemini
	ctx := context.Background()
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  apiKey,
		Backend: genai.BackendGeminiAPI,
	})

	model := "gemini-1.5-flash-002" // Ou "gemini-1.5-pro-002" para melhor qualidade
	result, err := client.Models.GenerateContent(ctx, model, genai.Text(prompt), nil)
	if err != nil {
		return "", err
	}

	// Exibir resposta
	return result.Text()
}
