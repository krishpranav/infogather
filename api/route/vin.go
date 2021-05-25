package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/krishpranav/infogather/api/handlers"
)

func Vin(app *fiber.App) {
	var h handlers.VinHandler
	r := app.Group("/vin")
	r.Get("/", h.Index)
	r.Get("/:vin", h.Show)
}
