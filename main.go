package main

import (
	telegram_router "telegram-bot/src/api/routers/telegram"
	"telegram-bot/src/config/env"
	"telegram-bot/src/worker"

	"github.com/gofiber/fiber/v2"
)

func main() {
	env.Start()
	go worker.Start()

	app := fiber.New()

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	telegram_router.SetupRoutes(app)

	app.Listen(":8080")
}
