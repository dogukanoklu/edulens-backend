package middleware

import "github.com/gofiber/fiber/v2"

func CorsMiddleware(c *fiber.Ctx) error {
    c.Set("Access-Control-Allow-Origin", "http://localhost:3000")
    c.Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT")
    c.Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With")

    if c.Method() == fiber.MethodOptions {
        return c.SendStatus(fiber.StatusNoContent)
    }

    return c.Next()
}
