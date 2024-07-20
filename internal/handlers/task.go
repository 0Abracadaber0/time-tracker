package handlers

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"log/slog"
	"time-tracker/internal/database"
	"time-tracker/internal/models"
)

func CreateTaskHandler(c *fiber.Ctx) error {
	log := c.Locals("logger").(*slog.Logger)

	var task models.Task
	if err := json.Unmarshal(c.Body(), &task); err != nil {
		log.Error("Invalid JSON format", "statusCode", 400)
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid JSON format",
		})
	}

	_, err := database.DB.Exec("INSERT INTO tasks (user_id, name, time) VALUES ($1, $2, $3)",
		c.Params("userId"),
		task.Name,
		0,
	)
	if err != nil {
		log.Error("Error inserting task", "user_id", c.Params("userId"))
	}

	log.Info("Task has been added to the database")

	return c.JSON(task)

}
