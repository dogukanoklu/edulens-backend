package utils

import (
	"fmt"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

func ParseAndValidate(c *fiber.Ctx, req interface{}) error {
	if err := c.BodyParser(req); err != nil {
		return err
	}

	if err := validate.Struct(req); err != nil {
		return err
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
