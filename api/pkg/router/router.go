package router


import (
	"api/pkg/handlers"

	"github.com/gofiber/fiber/v2"

)

func Router(app *fiber.App) {
	
	app.Get("/classes", handlers.GetClasses)
	app.Get("/attendance/classID", handlers.GetAttendanceWithStudents)
	

}
