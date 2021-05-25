package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/krishpranav/infogather/api/handlers"
)

func Username(app *fiber.App) {
	var h handlers.UsernameHandler
	r := app.Group("/username")
	r.Get("/", h.Index)
	r.Get("/:username", h.Show)
}
