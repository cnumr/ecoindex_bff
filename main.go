package main

import (
	"github.com/cnumr/ecoindex-bff/config"
	"github.com/cnumr/ecoindex-bff/handler"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

var ENV *config.Environment = config.GetEnvironment()

func main() {
	config.ENV = ENV

	app := fiber.New()

	app.Use(compress.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
	}))

	app.Get("/badge", handler.GetEcoindexBadge)
	app.Get("/redirect", handler.GetEcoindexRedirect)
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	api := app.Group("/api")
	api.Get("/results", handler.GetEcoindexResultsApi)
	api.Post("/tasks", handler.CreateTask)
	api.Get("/tasks/:id", handler.GetTask)
	api.Get("/screenshot/:id", handler.GetScreenshotApi)

	// Deprecated -->
	app.Get("/", handler.GetEcoindexResults)
	app.Post("/tasks", handler.CreateTask)
	app.Get("/tasks/:id", handler.GetTask)
	app.Get("/screenshot/:id", handler.GetScreenshotApi)
	// <-- Deprecated

	app.Listen(":" + config.ENV.AppPort)
}
