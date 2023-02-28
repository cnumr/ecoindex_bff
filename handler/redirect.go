package handler

import (
	"github.com/cnumr/ecoindex-bff/config"
	"github.com/cnumr/ecoindex-bff/services"
	"github.com/gofiber/fiber/v2"
)

func GetEcoindexRedirect(c *fiber.Ctx) error {
	_, ecoindexResults, shouldReturn, returnValue := services.HandleEcoindexRequest(c)
	if shouldReturn {
		return returnValue
	}

	if ecoindexResults.LatestResult.Id == "" {
		return c.Redirect(config.ENV.EcoindexUrl, fiber.StatusSeeOther)
	}

	return c.Redirect(config.ENV.EcoindexUrl+"/resultat/?id="+ecoindexResults.LatestResult.Id, fiber.StatusSeeOther)
}
