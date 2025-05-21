package router

import (
	"api/pkg/handlers"

	"github.com/gofiber/fiber/v2"
)

func Router(app *fiber.App) {

	app.Get("/v1/classes", handlers.GetClasses)

	app.Post("/v1/attendance/:classID", handlers.AddAttendance)
	app.Get("/v1/attendance/:classID", handlers.GetAttendanceWithStudents)
	app.Put("/v1/attendance/:attendanceID", handlers.UpdateAttendance)
	app.Get("/images/*", handlers.Images)

}
