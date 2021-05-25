package handlers

import (
	"github.com/gofiber/fiber/v2"
	hits "github.com/krishpranav/infogather/pkg/noauth/vin"
)

type VinHandler struct {
}

func (h VinHandler) Index(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusNotFound).SendFile("./views/pages/vin.html")
}

func (h VinHandler) Show(ctx *fiber.Ctx) error {
	vin := ctx.Params("vin")

	vininfo, err := hits.VinLookup(vin)
	if err != nil {
		return ctx.JSON(fiber.Map{"message": "failure"})
	}

	return ctx.JSON(vininfo)
}
