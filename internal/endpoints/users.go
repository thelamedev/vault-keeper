package endpoints

import (
	"encoding/json"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/thelamedev/vault-keeper/internal/database"
	"github.com/thelamedev/vault-keeper/internal/entities"
)

func CreateUser(c *fiber.Ctx) error {
	log.Printf("Got Request: %s", c.Locals("requestid"))

	var user entities.User
	err := json.Unmarshal(c.Body(), &user)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	err = database.DB.Create(&user).Error
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Failed to add user",
		})
	}

	return c.JSON(fiber.Map{
		"message": "User Created Successfully",
	})
}

func GetUser(c *fiber.Ctx) error {
	userID := c.Params("id")
	log.Printf("Got Request: %s | userId: %v", c.Locals("requestid"), userID)

	var user entities.User

	err := database.DB.First(&user, userID).Error
	if err != nil {
		if err.Error() == "record not found" {
			return c.Status(404).JSON(fiber.Map{
				"error": "User not found",
			})
		}
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to retrieve user",
		})
	}

	return c.JSON(user)
}
