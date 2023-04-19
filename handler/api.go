package handler

import (
	"fmt"

	"github.com/cnumr/ecoindex-bff/config"
	"github.com/cnumr/ecoindex-bff/services"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/proxy"
)

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
	proxy.Do(c, config.ENV.ApiUrl+"/v1/ecoindexes/"+c.Params("id")+"/screenshot")

	return nil
}

func ComputeEcoindex(c *fiber.Ctx) error {

	size := fmt.Sprintf("%f", c.QueryFloat("size"))
	dom := fmt.Sprintf("%d", c.QueryInt("dom"))
	requests := fmt.Sprintf("%d", c.QueryInt("requests"))

	c.Request().Header.Set("x-rapidapi-key", config.ENV.ApiKey)

	proxy.Do(c, config.ENV.ApiUrl+"/ecoindex"+"?dom="+dom+"&requests="+requests+"&size="+size)

	return nil
}
