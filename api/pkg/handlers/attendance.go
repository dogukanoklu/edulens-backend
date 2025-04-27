package handlers

import (
	"api/pkg/database"
	"api/pkg/models"
	"api/pkg/repository"
	"api/pkg/services"
	"api/pkg/utils"

	"github.com/gofiber/fiber/v2"
)

func AddAttendance(c *fiber.Ctx) error {
	classID := c.Params("classID")

	var req []models.ReqAddAttendance

	if err := utils.ParseAndValidate(c, &req); err != nil {
		return utils.JSONError(c, fiber.StatusBadRequest, "Invalid request payload.")
	}

	db := database.GetDB()
	attendanceRepo := repository.NewGormAttendanceRepository(db)
	attendanceService := services.NewAttendanceService(attendanceRepo)

	err := attendanceService.AddAttendance(classID, req)
	if err != nil {
		return utils.JSONError(c, fiber.StatusInternalServerError, "Failed to add attendance.")
	}

	return utils.JSONResponse(c, fiber.StatusNoContent, nil)

}

func GetAttendanceWithStudents(c *fiber.Ctx) error {
	classID := c.Params("classID")

	date, err := utils.ParseDateQuery(c)
	if err != nil {
		return utils.JSONError(c, fiber.StatusBadRequest, "Invalid parameters")
	}

	db := database.GetDB()
	attendanceRepo := repository.NewGormAttendanceRepository(db)
	attendanceService := services.NewAttendanceService(attendanceRepo)

	attendances, err := attendanceService.GetAttendanceWithStudents(classID, date)
	if err != nil {
		return utils.JSONError(c, fiber.StatusInternalServerError, "Failed to retrieve attendance information.")
	}

	if len(attendances.Students) == 0 {
		return utils.JSONError(c, fiber.StatusNotFound, "The requested attendance data could not be found for the specified parameters.")
	}

	return utils.JSONResponse(c, fiber.StatusOK, attendances)
}

func UpdateAttendance(c *fiber.Ctx) error {
	attendanceID := c.Params("attendanceID")

	var req []models.ReqUpdateAttendance

	if err := utils.ParseAndValidate(c, &req); err != nil {
		return utils.JSONError(c, fiber.StatusBadRequest, "Invalid request payload.")
	}

	db := database.GetDB()
	attendanceRepo := repository.NewGormAttendanceRepository(db)
	attendanceService := services.NewAttendanceService(attendanceRepo)

	if err := attendanceService.UpdateAttendance(attendanceID, req); err != nil {
		return utils.JSONError(c, fiber.StatusInternalServerError, "Failed to update attendance.")
	}

	return utils.JSONResponse(c, fiber.StatusNoContent, "")
}
