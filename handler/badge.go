package handler

import (
	"context"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/cnumr/ecoindex-bff/config"
	"github.com/cnumr/ecoindex-bff/services"
	"github.com/go-redis/cache/v8"
	"github.com/gofiber/fiber/v2"
)

func GetEcoindexBadge(c *fiber.Ctx) error {
	var grade, theme, badgeSvg string

	if c.Query("theme") == "dark" {
		theme = "dark"
	} else {
		theme = "light"
	}

	queryUrl, ecoindexResults, shouldReturn, returnValue := services.HandleEcoindexRequest(c)
	if shouldReturn {
		return returnValue
	}

	if ecoindexResults.LatestResult.Grade == "" {
		grade = "unkown"
	} else {
		grade = ecoindexResults.LatestResult.Grade
	}

	ctx := context.Background()
	cacheKey := "badge-" + grade + "-" + theme

	if err := config.CACHE.Get(ctx, cacheKey, &badgeSvg); err == nil {
		return sendBadgeSvg(c, queryUrl, badgeSvg, theme)
	}

	badgeSvg = services.GetBadgeSvg(grade, theme)
	if err := config.CACHE.Set(&cache.Item{
		Ctx:   ctx,
		Key:   cacheKey,
		Value: badgeSvg,
		TTL:   time.Duration(config.ENV.CacheTtl) * time.Minute,
	}); err != nil {
		log.Default().Println(err)
	}

	return sendBadgeSvg(c, queryUrl, badgeSvg, theme)
}

func sendBadgeSvg(c *fiber.Ctx, queryUrl string, badgeSvg string, theme string) error {
	c.Type("svg")
	c.Set("X-Ecoindex-Url", queryUrl)
	c.Set("X-Ecoindex-Theme", theme)
	c.Set(fiber.HeaderCacheControl, "public, max-age="+strconv.Itoa(config.ENV.CacheTtl))
	c.Set(fiber.HeaderLastModified, time.Now().Format(http.TimeFormat))
	c.Vary("X-Ecoindex-Url")
	c.Vary("X-Ecoindex-Theme")

	return c.SendString(badgeSvg)
}
