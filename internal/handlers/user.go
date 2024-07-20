package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log/slog"
	"strings"
	"time-tracker/config"
	"time-tracker/internal/database"
	"time-tracker/internal/models"
)

type PassportData struct {
	PassportSerie  string `json:"passportSerie"`
	PassportNumber string `json:"passportNumber"`
}

func CreateUserHandler(c *fiber.Ctx) error {
	log := c.Locals("logger").(*slog.Logger)

	var data PassportData
	if err := json.Unmarshal(c.Body(), &data); err != nil {
		log.Error("Invalid JSON format", "statusCode", 400)
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid JSON format",
		})
	}

	data.PassportSerie, data.PassportNumber =
		strings.Split(data.PassportNumber, " ")[0],
		strings.Split(data.PassportNumber, " ")[1]

	url := config.Config("URL_API")

	a := fiber.AcquireAgent()
	req := a.Request()
	req.Header.SetMethod("GET")
	params := fmt.Sprintf("/info?passportNumber=%s&passportSerie=%s", data.PassportNumber, data.PassportSerie)
	req.SetRequestURI(url + params)

	if err := a.Parse(); err != nil {
		log.Error("Error parsing request", "url", url, "statusCode", 400)
	}
	log.Info("Request has been sent", "url", url+params)

	statusCode, body, errs := a.Bytes()
	if errs != nil {
		log.Error("Error parsing request", "url", url, "statusCode", 500)
		return c.Status(500).JSON(fiber.Map{
			"error": errs,
		})
	}

	if statusCode != 200 {
		log.Error("Error parsing request", "url", url, "statusCode", statusCode)
		return c.Status(statusCode).JSON(fiber.Map{
			"error": "Invalid JSON format",
		})
	}

	var response models.User
	err := json.Unmarshal(body, &response)
	if err != nil {
		log.Error("Error parsing response", "url", url, "statusCode", 400)
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid JSON format",
		})
	}

	log.Info("Response has been received", "url", url+params, "statusCode", statusCode)

	_, err = database.DB.Exec("INSERT INTO users (surname, name, patronymic, address) VALUES ($1, $2, $3, $4)",
		response.Surname,
		response.Name,
		response.Patronymic,
		response.Address,
	)

	if err != nil {
		log.Error(err.Error())
	}

	log.Info("User has been added to the database")

	return c.JSON(response)
}
