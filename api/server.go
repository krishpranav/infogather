package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	"github.com/krishpranav/infogather/api/handlers"
	"github.com/krishpranav/infogather/api/route"
)

// function for launching the server
func main() {
	app := fiber.New(fiber.Config{
		// Pass view engine
		Views: html.New("./views", ".html"),
		// Pass global error handler
		ErrorHandler: handlers.Errors("./public/500.html"),
	})

	// Render index template with IP input value
	app.Get("/", handlers.Render())

	// Serve static assets
	app.Static("/", "./public", fiber.Static{
		Compress: true,
	})

	// Main GEO handler that is cached for 10 minutes
	route.Username(app)
	route.Discord(app)
	route.Ip(app)
	route.Vin(app)

	// Handle 404 errors
	app.Use(handlers.NotFound("./public/404.html"))

	// Listen on port :3000
	log.Fatal(app.Listen(":6174"))
}
