package utils

import (
	"api/pkg/log"
	"api/pkg/models"
	"fmt"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

func ParseAndValidate(c *fiber.Ctx, req interface{}) error {
	// Gelen raw veriyi logla
	log.Info("Raw body data:"+ string(c.Body())) // Gelen raw body'yi logla

	// Gelen veriyi dizi olarak parse et
	if err := c.BodyParser(req); err != nil {
		log.Error("Error parsing body:", err)
		return err
	}


	// Gelen verinin türünü kontrol et ve her bir elemanı validasyonla kontrol et
	switch v := req.(type) {
	case *[]models.ReqUpdateAttendance:
		// Eğer req bir []models.ReqUpdateAttendance dizisi ise
		for i := 0; i < len(*v); i++ {
			// Her bir elemanı validasyonla kontrol et
			if err := validate.Struct((*v)[i]); err != nil {
				log.Error("Validation failed for element:", err)
				return err
			}
		}
	default:
		// Eğer farklı bir tip ise, validasyonu uygulamak
		if err := validate.Struct(req); err != nil {
			log.Error("Validation failed:", err)
			return err
		}
	}

	return nil
}

func ParseDateQuery(c *fiber.Ctx) (int64, error) {
	dateStr := c.Query("date", fmt.Sprintf("%d", time.Now().Unix()))
	date, err := strconv.ParseInt(dateStr, 10, 64)
	if err != nil {
		return 0, err
	}
	return date, nil
}
