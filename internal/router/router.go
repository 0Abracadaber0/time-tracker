package router

import (
	"github.com/gofiber/fiber/v2"
	"time-tracker/internal/handlers"
)

func SetupRoutes(app *fiber.App) {
	// Return list of users
	app.Get("/users", handlers.ShowAllUsersHandler)
	// Return all tasks of user
	app.Get("/users/:userId/tasks", handlers.ShowAllTasksHandler)

	// Start timer on the users task
	app.Post("/tasks/:taskId/start", handlers.StartTimerHandler)
	// Stop timer on the users task
	app.Post("tasks/:taskId/stop", handlers.StopTimerHandler)

	// Delete the user
	app.Delete("/users/:userId", func(c *fiber.Ctx) error { return nil })

	// Update the user
	app.Put("/users/:userId/", func(c *fiber.Ctx) error { return nil })

	// Create new user
	app.Post("/users", handlers.CreateUserHandler)

	// Create new task
	app.Post("/users/:userId/tasks", handlers.CreateTaskHandler)

}
