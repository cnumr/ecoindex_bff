package main

import (
	"github.com/cnumr/ecoindex-bff/config"
	"github.com/cnumr/ecoindex-bff/handler"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

var ENV *config.Environment = config.GetEnvironment()

func main() {
	config.ENV = ENV

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
	}))

	app.Get("/", handler.GetEcoindexBadge)
	app.Post("/tasks", handler.CreateTask)
	app.Get("/tasks/:id", handler.GetTask)
	app.Get("/screenshot/:id", handler.GetScreenshot)

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	app.Listen(":" + config.ENV.AppPort)
}
