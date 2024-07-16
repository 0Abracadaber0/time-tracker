package router

import (
	"github.com/gofiber/fiber/v2"
	"time-tracker/internal/handlers"
)

func SetupRoutes(app *fiber.App) {
	// Return list of users
	app.Get("/users", func(c *fiber.Ctx) error { return nil })
	// Return all tasks of user
	app.Get("/users/:userId/tasks", func(c *fiber.Ctx) error { return nil })

	// Start timer on the users task
	app.Post("/users/:userId/tasks/:taskId/start", func(c *fiber.Ctx) error { return nil })
	// Stop timer on the users task
	app.Post("/users/:userId/tasks/:taskId/stop", func(c *fiber.Ctx) error { return nil })

	// Delete the user
	app.Delete("/users/:userId", func(c *fiber.Ctx) error { return nil })

	// Update the user
	app.Put("/users/:userId/", func(c *fiber.Ctx) error { return nil })

	// Create new user
	app.Post("/users", handlers.CreateUserHandler)

}
