package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func Errors(file string) fiber.ErrorHandler {
	return func(c *fiber.Ctx, err error) error {
		fmt.Println(err.Error())
		return c.Status(fiber.StatusInternalServerError).SendFile(file)
	}
}
