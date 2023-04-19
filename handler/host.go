package handler

import (
	"github.com/cnumr/ecoindex-bff/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/proxy"
)

func GetHost(c *fiber.Ctx) error {
	c.Request().Header.Set("x-rapidapi-key", config.ENV.ApiKey)
	proxy.Do(c, config.ENV.ApiUrl+"/v1/hosts/"+c.Params("hostname"))

	return nil
}
