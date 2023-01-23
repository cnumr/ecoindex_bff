package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/proxy"
	"github.com/vvatelot/ecoindex-bff/config"
)

func CreateTask(c *fiber.Ctx) error {
	c.Request().Header.Set("x-rapidapi-key", config.ENV.ApiKey)
	proxy.Forward(config.ENV.ApiUrl + "/v1/tasks/ecoindexes")(c)

	return nil
}

func GetTask(c *fiber.Ctx) error {
	c.Request().Header.Set("x-rapidapi-key", config.ENV.ApiKey)
	proxy.Forward(config.ENV.ApiUrl + "/v1/tasks/ecoindexes/" + c.Params("id"))(c)

	return nil
}
