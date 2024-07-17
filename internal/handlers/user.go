package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"strings"
	"time-tracker/config"
	"time-tracker/internal/models"
)

type PassportData struct {
	PassportSerie  string `json:"passportSerie"`
	PassportNumber string `json:"passportNumber"`
}

func CreateUserHandler(c *fiber.Ctx) error {

	var data PassportData
	if err := json.Unmarshal(c.Body(), &data); err != nil {
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

	fmt.Println(req.String())

	if err := a.Parse(); err != nil {
		panic(err)
	}

	statusCode, body, errs := a.Bytes()
	if errs != nil {
		return c.Status(500).JSON(fiber.Map{"error": errs})
	}

	if statusCode != 200 {
		return c.Status(statusCode).JSON(fiber.Map{
			"error": "Invalid JSON format",
		})
	}

	var response models.User
	err := json.Unmarshal(body, &response)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid JSON format",
		})
	}

	return c.JSON(response)
}
