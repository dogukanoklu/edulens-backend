package utils

import (
	"github.com/gofiber/fiber/v2"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

func JSONError(c *fiber.Ctx, status int, message string) error {
	c.Status(status)
	return c.JSON(ErrorResponse{Error: message})
}

func JSONResponse(c *fiber.Ctx, status int, data interface{}) error {
	c.Status(status)
	return c.JSON(data)
}
