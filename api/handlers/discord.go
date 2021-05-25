package handlers

import (
	"github.com/gofiber/fiber/v2"
)

type DiscordHandler struct {
}

func (h DiscordHandler) Index(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusNotFound).SendFile("./views/pages/discord.html")
}
