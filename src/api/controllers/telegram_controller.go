package controllers

import (
	"github.com/gofiber/fiber/v2"
)

type TelegramController struct{}

func NewTelegramController() *TelegramController {
	return &TelegramController{}
}

func (t *TelegramController) Send(c *fiber.Ctx) error {
	return c.SendString("Message sent")
}
