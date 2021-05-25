package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/krishpranav/infogather/api/handlers"
)

// ip route
func Ip(app *fiber.App) {
	var h handlers.IpHandler
	r := app.Group("/ip")
	r.Get("/", h.Index)
	r.Get("/:ip", h.Show)
}
