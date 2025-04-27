package main

import (
	"api/pkg/database"
	"api/pkg/log"
	"api/pkg/middleware"
	"api/pkg/router"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Error("Error loading the .env file: %v", err)
	}

	database.InitDB()

	app := fiber.New()
	app.Use(middleware.CorsMiddleware)
	router.Router(app)

	log.Info("Server listenin on http://localhost:8000 :)")

	if err := app.Listen(":8000"); err != nil {
		log.Error("There was an error with the http server: %v", err)
	}

}
