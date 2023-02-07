package handler

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/cnumr/ecoindex-bff/assets"
	"github.com/cnumr/ecoindex-bff/config"
	"github.com/cnumr/ecoindex-bff/models"
	"github.com/cnumr/ecoindex-bff/services"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/proxy"
)

var badgeTemplate *template.Template

// Deprecated
func GetEcoindexResults(c *fiber.Ctx) error {
	queryUrl, ecoindexResults, shouldReturn, returnValue := handleEcoindexRequest(c)
	if shouldReturn {
		return returnValue
	}

	if c.Query("badge") == "true" {
		c.Type("svg")
		c.Set("X-Ecoindex-Url", queryUrl)
		c.Set(fiber.HeaderCacheControl, "public, max-age="+config.ENV.CacheControl)
		c.Set(fiber.HeaderLastModified, time.Now().Format(http.TimeFormat))
		c.Vary("X-Ecoindex-Url")
		return c.SendString(generateBadge(ecoindexResults))
	}

	if ecoindexResults.Count == 0 {
		c.Status(fiber.ErrNotFound.Code)
	}

	return c.JSON(ecoindexResults)
}

func GetEcoindexBadgeJs(c *fiber.Ctx) error {
	mediaType := "application/javascript"
	c.Type("js")
	c.Set(fiber.HeaderCacheControl, "public, max-age="+config.ENV.CacheControl)
	c.Set(fiber.HeaderLastModified, time.Now().Format(http.TimeFormat))

	input, err := os.ReadFile("./assets/js/badge.js")
	if err != nil {
		panic(err)
	}

	javascript := bytes.Replace(input, []byte("{{url}}"), []byte(config.ENV.AppUrl), -1)

	js := minifyString(mediaType, string(javascript))

	return c.SendString(js)
}

func GetEcoindexRedirect(c *fiber.Ctx) error {
	_, ecoindexResults, shouldReturn, returnValue := handleEcoindexRequest(c)
	if shouldReturn {
		return returnValue
	}

	if ecoindexResults.Count == 0 {
		return c.Redirect(config.ENV.EcoindexUrl, fiber.StatusSeeOther)
	}

	return c.Redirect(config.ENV.EcoindexUrl+"/resultat/?id="+ecoindexResults.LatestResult.Id, fiber.StatusSeeOther)
}

func GetEcoindexResultsApi(c *fiber.Ctx) error {
	_, ecoindexResults, shouldReturn, returnValue := handleEcoindexRequest(c)
	if shouldReturn {
		return returnValue
	}

	if ecoindexResults.Count == 0 {
		c.Status(fiber.ErrNotFound.Code)
	}

	return c.JSON(ecoindexResults)
}

func GetEcoindexBadge(c *fiber.Ctx) error {
	queryUrl, ecoindexResults, shouldReturn, returnValue := handleEcoindexRequest(c)
	if shouldReturn {
		return returnValue
	}

	c.Type("svg")
	c.Set("X-Ecoindex-Url", queryUrl)
	c.Set(fiber.HeaderCacheControl, "public, max-age="+config.ENV.CacheControl)
	c.Set(fiber.HeaderLastModified, time.Now().Format(http.TimeFormat))
	c.Vary("X-Ecoindex-Url")

	return c.SendString(generateBadge(ecoindexResults))
}

func GetScreenshotApi(c *fiber.Ctx) error {
	c.Request().Header.Set("x-rapidapi-key", config.ENV.ApiKey)
	proxy.Forward(config.ENV.ApiUrl + "/v1/ecoindexes/" + c.Params("id") + "/screenshot")(c)

	return nil
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

	return minifyString("image/svg+xml", buf.String())
}

func handleEcoindexRequest(c *fiber.Ctx) (string, models.EcoindexSearchResults, bool, error) {
	queryUrl := c.Query("url")

	urlToAnalyze, err := url.ParseRequestURI(queryUrl)
	if err != nil || urlToAnalyze.Host == "" {
		c.Status(fiber.ErrBadRequest.Code)

		return "", models.EcoindexSearchResults{}, true, c.SendString("Url to analyze is invalid")
	}

	ecoindexResults, err := services.GetEcoindexResults(urlToAnalyze.Host, urlToAnalyze.Path)
	if err != nil {
		panic(err)
	}

	return queryUrl, ecoindexResults, false, nil
}

func minifyString(mediaType string, input string) string {
	minified, err := config.MINIFIER.String(mediaType, input)
	if err != nil {
		panic(err)
	}

	return minified
}
