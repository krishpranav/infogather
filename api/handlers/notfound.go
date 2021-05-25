package handlers

import "github.com/gofiber/fiber/v2"

func NotFound(file string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).SendFile(file)
	}
}
