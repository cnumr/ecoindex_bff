package handler

import (
	"github.com/cnumr/ecoindex-bff/config"
	"github.com/cnumr/ecoindex-bff/services"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/proxy"
)

// Deprecated
func GetEcoindexResults(c *fiber.Ctx) error {
	if c.Query("badge") == "true" {
		return GetEcoindexBadge(c)
	}

	return GetEcoindexResultsApi(c)
}

func GetEcoindexResultsApi(c *fiber.Ctx) error {
	_, ecoindexResults, shouldReturn, returnValue := services.HandleEcoindexRequest(c)
	if shouldReturn {
		return returnValue
	}

	if ecoindexResults.Count == 0 {
		c.Status(fiber.ErrNotFound.Code)
	}

	return c.JSON(ecoindexResults)
}

func GetScreenshotApi(c *fiber.Ctx) error {
	c.Request().Header.Set("x-rapidapi-key", config.ENV.ApiKey)
	proxy.Forward(config.ENV.ApiUrl + "/v1/ecoindexes/" + c.Params("id") + "/screenshot")(c)

	return nil
}
