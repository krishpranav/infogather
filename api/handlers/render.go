package handlers

import "github.com/gofiber/fiber/v2"

func Render() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"IP": c.IP(),
		})
	}
}
