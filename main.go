package main

import (
	telegram_router "telegram-bot/src/api/routers/telegram"
	"telegram-bot/src/config/client"
	"telegram-bot/src/config/db"
	"telegram-bot/src/config/env"
	"telegram-bot/src/worker"

	"github.com/gofiber/fiber/v2"
)

func main() {
	env.Start()
	client.Start()
	db.ConnectMongoDB()
	go worker.Start()

	app := fiber.New()

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	telegram_router.SetupRoutes(app)

	app.Listen(":8080")
}
