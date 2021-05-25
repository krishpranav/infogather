package handlers

import "github.com/gofiber/fiber/v2"

type IpHandler struct {
}

func (h IpHandler) Index(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusNotFound).SendFile("./views/pages/ip.html")
}
