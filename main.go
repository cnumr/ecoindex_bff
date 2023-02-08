package main

import (
	"github.com/cnumr/ecoindex-bff/config"
	"github.com/cnumr/ecoindex-bff/handler"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/tdewolff/minify/v2"
)

var ENV *config.Environment = config.GetEnvironment()
var MINIFIER *minify.M = config.GetMinifier()

func main() {
	config.ENV = ENV
	config.MINIFIER = MINIFIER

	app := fiber.New()
	if ENV.Env == "dev" {
		app.Use(logger.New(logger.Config{
			Format: "[${time}] | ${status} | ${latency} | ${method} | ${path} | url=${query:url}\n",
		}))
	}
	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestCompression,
	}))
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
	}))

	app.Get("/js/badge.js", handler.GetEcoindexBadgeJs)
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
