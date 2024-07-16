package main

import (
	"github.com/gofiber/fiber/v2"
	"time-tracker/internal/router"
)

func main() {
	app := fiber.New()

	router.SetupRoutes(app)

	err := app.Listen("0.0.0.0:8080")
	if err != nil {
		return
	}
}
