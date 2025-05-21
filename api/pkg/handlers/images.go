package handlers

import (
	"path/filepath"
	"github.com/gofiber/fiber/v2"
)

// http://localhost:8000/images/1014.jpg
func Images(c *fiber.Ctx) error {
	// Parametreyi al (örneğin: 1014.jpg)
	imageName := c.Params("*") // joker parametreyle tüm yolu alır

	// Gerçek yol: ../../ai/data/student_images/1014.jpg
	imagePath := filepath.Join("..", "..", "ai", "data", "student_images", imageName)

	// Dosyayı gönder
	return c.SendFile(imagePath, true)
}
