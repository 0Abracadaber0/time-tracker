package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log/slog"
	"time"
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

	_, err := database.DB.Exec("INSERT INTO tasks (user_id, name, time, is_working) VALUES ($1, $2, $3, $4)",
		c.Params("userId"),
		task.Name,
		0,
		"false",
	)
	if err != nil {
		log.Error("Error inserting task", "user_id", c.Params("userId"))
	}

	log.Info("Task has been added to the database")

	return c.JSON(task)
}

func StartTimerHandler(c *fiber.Ctx) error {
	log := c.Locals("logger").(*slog.Logger)

	taskId := c.Params("taskId")
	if taskId == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "taskId is required"})
	}

	query := fmt.Sprintf("UPDATE tasks SET last_start = '%s', is_working = true WHERE user_id = '%s'",
		time.Now().Format("2006-01-02 15:04:05"),
		taskId,
	)

	_, err := database.DB.Exec(query)
	if err != nil {
		log.Error(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal server error"})
	}

	log.Info("Timer started for task", "task_id", c.Params("taskId"))
	return c.SendStatus(fiber.StatusOK)
}
