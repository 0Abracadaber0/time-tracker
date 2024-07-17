package main

import (
	"github.com/gofiber/fiber/v2"
	"time-tracker/config"
	"time-tracker/internal/router"
)

func main() {
	log := config.SetupLogger(config.Config("ENV"))
	log.Info("Starting application")

	app := fiber.New()

	app.Use(func(ctx *fiber.Ctx) error {
		ctx.Locals("logger", log)
		return ctx.Next()
	})

	router.SetupRoutes(app)

	err := app.Listen("0.0.0.0:8080")
	if err != nil {
		log.Error(err.Error())
	}
}
