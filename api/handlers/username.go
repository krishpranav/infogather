package handlers

import "github.com/gofiber/fiber/v2"

type UsernameHandler struct {
}

func (h UsernameHandler) Index(ctx *fiber.Ctx) {
	return ctx.Status(fiber.StatusNotFound).SendFile("./views/pages/username.html")
}
