package telegram_router

import (
	"telegram-bot/src/api/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	controller := controllers.NewTelegramController()

	app.Get("/", controller.Send)
}
