package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vvatelot/ecoindex-microfront/config"
	"github.com/vvatelot/ecoindex-microfront/services"
)

func GetEcoindexBadge(c *fiber.Ctx) error {
	var color, grade string

	ecoindexUrl := config.ENV.EcoindexUrl

	ecoindex, err := services.GetEcoindex(c.Query("url"))
	if err != nil {
		panic(err)
	}

	if ecoindex.Id == "" {
		color = "light-grey"
		grade = "?"
	} else {
		color = services.GetColor(ecoindex.Grade)
		grade = ecoindex.Grade
		ecoindexUrl = ecoindexUrl + "/resultat/?id=" + ecoindex.Id
	}

	return c.Render("badge", fiber.Map{
		"Grade": grade,
		"Id":    ecoindex.Id,
		"Color": color,
		"Score": ecoindex.Score,
		"Date":  ecoindex.Date,
		"Url":   ecoindexUrl,
	})
}
