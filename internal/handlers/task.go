package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log/slog"
	"strconv"
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

	query := fmt.Sprintf("UPDATE tasks SET last_start = '%s', is_working = true WHERE task_id = '%s'",
		time.Now().UTC().Format("2006-01-02 15:04:05"),
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

func StopTimerHandler(c *fiber.Ctx) error {
	log := c.Locals("logger").(*slog.Logger)
	taskId := c.Params("taskId")
	if taskId == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "taskId is required"})
	}

	var startTime time.Time
	var minutes int
	query := fmt.Sprintf("SELECT last_start, time FROM tasks WHERE task_id = '%s'",
		c.Params("taskId"),
	)
	row := database.DB.QueryRow(query)
	err := row.Scan(&startTime, &minutes)
	if err != nil {
		log.Error(err.Error())
	}

	now := time.Now().UTC()
	diff := now.Sub(startTime)

	log.Debug("", "diff", diff)

	minutes += int(diff.Minutes())
	log.Debug(strconv.Itoa(minutes))

	query = fmt.Sprintf("UPDATE tasks SET is_working = false, time = %d WHERE task_id = '%s'",
		minutes,
		taskId,
	)
	_, err = database.DB.Exec(query)
	if err != nil {
		log.Error(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal server error"})
	}
	log.Info("Timer stopped for task", "task_id", c.Params("taskId"))
	return c.SendStatus(fiber.StatusOK)
}
