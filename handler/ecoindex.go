package handler

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/vvatelot/ecoindex-microfront/config"
	"github.com/vvatelot/ecoindex-microfront/services"
)

func GetEcoindexBadge(c *fiber.Ctx) error {
	var color, grade, title, score string
	queryUrl := c.Query("url")
	if queryUrl == "" {
		return c.SendStatus(fiber.ErrBadRequest.Code)
	}

	ecoindexUrl := config.ENV.EcoindexUrl

	ecoindex, err := services.GetEcoindex(queryUrl)
	if err != nil {
		panic(err)
	}

	if ecoindex.Id == "" {
		color = "light-grey"
		grade = "?"
		title = "Aucun résultat pour " + queryUrl
	} else {
		color = services.GetColor(ecoindex.Grade)
		grade = ecoindex.Grade
		ecoindexUrl = ecoindexUrl + "/resultat/?id=" + ecoindex.Id
		score = fmt.Sprintf("%f", ecoindex.Score)
		title = score + " / 100 au " + ecoindex.Date
	}

	return c.Render("badge", fiber.Map{
		"Grade": grade,
		"Id":    ecoindex.Id,
		"Color": color,
		"Title": title,
		"Date":  ecoindex.Date,
		"Url":   ecoindexUrl,
	})
}
