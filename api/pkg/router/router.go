package router


import (
	"api/pkg/handlers"

	"github.com/gofiber/fiber/v2"

)

func Router(app *fiber.App) {
	
	app.Get("/v1/classes", handlers.GetClasses)
	app.Get("/v1/attendance/classID", handlers.GetAttendanceWithStudents)
	

}
