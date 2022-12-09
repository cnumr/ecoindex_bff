package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/vvatelot/ecoindex-microfront/config"
	"github.com/vvatelot/ecoindex-microfront/handler"
)

var ENV *config.Environment = config.GetEnvironment()

func main() {
	config.ENV = ENV

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
	}))

	app.Get("/", handler.GetEcoindexBadge)

	app.Listen(":" + config.ENV.AppPort)
}
