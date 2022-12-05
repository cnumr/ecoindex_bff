package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vvatelot/ecoindex-microfront/services"
)

func GetEcoindexBadge(c *fiber.Ctx) error {
	ecoindex, err := services.GetEcoindex(c.Query("url"))
	if err != nil {
		panic(err)
	}

	if ecoindex.Id == "" {
		return c.SendString("")
	}

	return c.Render("badge", fiber.Map{
		"Grade": ecoindex.Grade,
		"Id":    ecoindex.Id,
		"Color": services.GetColor(ecoindex.Grade),
		"Score": ecoindex.Score,
		"Date":  ecoindex.Date,
	})
}
