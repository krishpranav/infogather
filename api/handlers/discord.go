package handlers

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/krishpranav/infogather/pkg/noauth/discord"
)

type DiscordHandler struct {
}

func (h DiscordHandler) Index(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusNotFound).SendFile("./views/pages/discord.html")
}

func (h DiscordHandler) Show(ctx *fiber.Ctx) error {
	token := ctx.Params("token")

	tokeninfo, err := discord.TokenLookup(token)
	if err != nil {
		return errors.New("error")
	}

	return ctx.JSON(tokeninfo)
}
