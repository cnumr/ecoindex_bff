package handler

import (
	"bytes"
	"fmt"
	"html/template"
	"net/url"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/proxy"
	"github.com/vvatelot/ecoindex-bff/assets"
	"github.com/vvatelot/ecoindex-bff/config"
	"github.com/vvatelot/ecoindex-bff/models"
	"github.com/vvatelot/ecoindex-bff/services"
)

var badgeTemplate *template.Template

func GetEcoindexBadge(c *fiber.Ctx) error {
	queryUrl := c.Query("url")

	urlToAnalyze, err := url.ParseRequestURI(queryUrl)
	if err != nil || urlToAnalyze.Host == "" {
		c.Status(fiber.ErrBadRequest.Code)
		return c.SendString("Url to analyze is invalid")
	}

	ecoindexResults, err := services.GetEcoindexResults(urlToAnalyze.Host, urlToAnalyze.Path)
	if err != nil {
		panic(err)
	}

	if c.Query("badge") == "true" {
		c.Type("svg")
		return c.SendString(generateBadge(ecoindexResults))
	}

	return c.JSON(ecoindexResults)
}

func initTemplate() {
	badgeTemplate = template.Must(template.ParseFS(&assets.TemplateFs, "template/badge.svg"))
}

func generateBadge(result models.EcoindexSearchResults) string {
	initTemplate()
	var color, grade, title, score string
	ecoindexUrl := config.ENV.EcoindexUrl

	if result.LatestResult.Id == "" {
		color = "light-grey"
		grade = "?"
		title = "Aucun résultat"
	} else {
		color = services.GetColor(result.LatestResult.Grade)
		grade = result.LatestResult.Grade
		ecoindexUrl = ecoindexUrl + "/resultat/?id=" + result.LatestResult.Id
		score = fmt.Sprintf("%.2f", result.LatestResult.Score)
		title = score + " / 100 au " + result.LatestResult.Date
	}

	vars := fiber.Map{
		"Grade": grade,
		"Id":    result.LatestResult.Id,
		"Color": color,
		"Title": title,
		"Date":  result.LatestResult.Date,
		"Url":   ecoindexUrl,
	}

	buf := &bytes.Buffer{}
	badgeTemplate.Execute(buf, vars)

	return buf.String()
}

func GetScreenshot(c *fiber.Ctx) error {
	c.Request().Header.Set("x-rapidapi-key", config.ENV.ApiKey)
	proxy.Forward(config.ENV.ApiUrl + "/v1/ecoindexes/" + c.Params("id") + "/screenshot")(c)

	return nil
}
