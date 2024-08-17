package internal

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"

	"github.com/thelamedev/vault-keeper/internal/config"
	"github.com/thelamedev/vault-keeper/internal/database"
	"github.com/thelamedev/vault-keeper/internal/endpoints"
	"github.com/thelamedev/vault-keeper/internal/entities"
)

func LoadServer(config *config.Config) *fiber.App {
	app := fiber.New(fiber.Config{
		AppName: "VaultKeeper v0.1",
	})

	database.DB.AutoMigrate(&entities.User{}, &entities.Project{}, &entities.APIKey{}, &entities.Secret{})

	LoadMiddlewares(app)
	LoadEndpoints(app, "/api/v1")

	return app
}

func LoadMiddlewares(app *fiber.App) {
	app.Use(requestid.New())
	app.Use(logger.New())
}

func LoadEndpoints(app *fiber.App, prefix string) {
	router := app.Group(prefix)

	router.Post("/users/add", endpoints.CreateUser)
	router.Get("/users/i/:id", endpoints.GetUser)
}
