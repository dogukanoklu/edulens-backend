package handlers

import (
	"api/pkg/database"
	"api/pkg/repository"
	"api/pkg/services"
	"api/pkg/utils"

	"github.com/gofiber/fiber/v2"
)

func GetClasses(c *fiber.Ctx) error {
	db := database.GetDB()
	classRepo := repository.NewGormClassRepository(db)
	classService := services.NewClassService(classRepo)

	classs, err := classService.GetClasses()
	if err != nil {
		return utils.JSONError(c, fiber.StatusInternalServerError, "Failed to retrieve class information.")
	}

	return utils.JSONResponse(c, fiber.StatusOK, classs)
}
