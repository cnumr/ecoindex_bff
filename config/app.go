package config

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func ConfigureApp(app *fiber.App) {
	if ENV.Env == "dev" {
		app.Use(logger.New(logger.Config{
			Format: "[${time}] | ${status} | ${latency} | ${method} | ${path} | url=${query:url} | refresh=${query:refresh} \n",
		}))
	}

	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestCompression,
	}))

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
	}))
}
