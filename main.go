package main

import (
	"github.com/cnumr/ecoindex-bff/config"
	"github.com/cnumr/ecoindex-bff/handler"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
)

func main() {
	config.ENV = config.GetEnvironment()
	config.CACHE = config.GetCache()
	config.MINIFIER = config.GetMinifier()

	app := *fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})
	config.ConfigureApp(&app)

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

	app.Listen(":" + config.ENV.AppPort)
}
