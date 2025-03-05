package client

import (
	"context"

	"telegram-bot/src/config/env"

	"google.golang.org/genai"
)

var (
	GlobalClient *genai.Client
	Ctx          = context.Background()
)

func Start() {
	var err error
	Ctx = context.Background()
	GlobalClient, err = genai.NewClient(Ctx, &genai.ClientConfig{
		APIKey:  env.GeminiApiToken,
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		panic(err)
	}
}
