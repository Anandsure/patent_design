package views

import (
	"github.com/gofiber/fiber/v2"
)

func search_es(c *fiber.Ctx) error {
	return c.SendString("search running")
}
