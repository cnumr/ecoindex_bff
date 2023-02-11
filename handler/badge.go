package handler

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"time"

	"github.com/cnumr/ecoindex-bff/assets"
	"github.com/cnumr/ecoindex-bff/config"
	"github.com/cnumr/ecoindex-bff/helper"
	"github.com/cnumr/ecoindex-bff/models"
	"github.com/cnumr/ecoindex-bff/services"
	"github.com/gofiber/fiber/v2"
)

func GetEcoindexBadgeJs(c *fiber.Ctx) error {
	var theme string

	mediaType := "application/javascript"
	c.Type("js")
	c.Set(fiber.HeaderCacheControl, "public, max-age="+strconv.Itoa(config.ENV.CacheTtl))
	c.Set(fiber.HeaderLastModified, time.Now().Format(http.TimeFormat))
	c.Set("X-Theme", "dark")
	c.Vary("X-Theme")

	input, err := assets.JsFs.ReadFile("js/badge.js")
	if err != nil {
		panic(err)
	}

	if c.Query("theme") == "dark" {
		theme = "dark"
		c.Set("X-Theme", "dark")
	} else {
		theme = "light"
		c.Set("X-Theme", "light")
	}

	javascript := bytes.Replace(input, []byte("{{url}}"), []byte(config.ENV.AppUrl), -1)
	javascript = bytes.Replace(javascript, []byte("{{theme}}"), []byte(theme), -1)

	js := helper.MinifyString(mediaType, string(javascript))

	return c.SendString(js)
}

func GetEcoindexBadge(c *fiber.Ctx) error {
	var dark bool

	queryUrl, ecoindexResults, shouldReturn, returnValue := services.HandleEcoindexRequest(c)
	if shouldReturn {
		return returnValue
	}

	if c.Query("theme") == "dark" {
		dark = true
	}

	c.Type("svg")
	c.Set("X-Ecoindex-Url", queryUrl)
	c.Set(fiber.HeaderCacheControl, "public, max-age="+strconv.Itoa(config.ENV.CacheTtl))
	c.Set(fiber.HeaderLastModified, time.Now().Format(http.TimeFormat))
	c.Vary("X-Ecoindex-Url")

	return c.SendString(generateBadge(ecoindexResults, dark))
}

func initTemplate(dark bool) {
	var badge string

	if dark {
		badge = "template/badge-dark.svg"
	} else {
		badge = "template/badge.svg"
	}
	badgeTemplate = template.Must(template.ParseFS(&assets.TemplateFs, badge))
}

func generateBadge(result models.EcoindexSearchResults, dark bool) string {
	initTemplate(dark)
	var color, grade, title, score string
	ecoindexUrl := config.ENV.EcoindexUrl

	if result.LatestResult.Id == "" {
		color = "grey"
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

	return helper.MinifyString("image/svg+xml", buf.String())
}
