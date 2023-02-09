package services

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/cnumr/ecoindex-bff/config"
	"github.com/cnumr/ecoindex-bff/models"
	"github.com/go-redis/cache/v8"
	"github.com/gofiber/fiber/v2"
)

func HandleEcoindexRequest(c *fiber.Ctx) (string, models.EcoindexSearchResults, bool, error) {
	ctx := context.Background()

	queryUrl := c.Query("url")

	if c.Query("refresh") != "true" {
		var wanted models.EcoindexSearchResults
		if err := config.CACHE.Get(ctx, queryUrl, &wanted); err == nil {
			return queryUrl, wanted, false, nil
		}
	}

	urlToAnalyze, err := url.ParseRequestURI(queryUrl)
	if err != nil || urlToAnalyze.Host == "" {
		c.Status(fiber.ErrBadRequest.Code)

		return "", models.EcoindexSearchResults{}, true, c.SendString("Url to analyze is invalid")
	}

	ecoindexResults, err := GetEcoindexResults(urlToAnalyze.Host, urlToAnalyze.Path)
	if err != nil {
		panic(err)
	}

	if err := config.CACHE.Set(&cache.Item{
		Ctx:   ctx,
		Key:   queryUrl,
		Value: ecoindexResults,
		TTL:   time.Duration(config.ENV.CacheTtl) * time.Minute,
	}); err != nil {
		panic(err)
	}

	return queryUrl, ecoindexResults, false, nil
}

func GetEcoindexResults(host string, path string) (models.EcoindexSearchResults, error) {
	client := http.Client{}
	req, err := http.NewRequest(http.MethodGet, config.ENV.ApiUrl+"/v1/ecoindexes", nil)
	if err != nil {
		return models.EcoindexSearchResults{}, err
	}

	q := req.URL.Query()
	q.Add("host", host)
	q.Add("size", "100")

	req.URL.RawQuery = q.Encode()
	req.Header = http.Header{
		"x-rapidapi-key": {config.ENV.ApiKey},
	}

	resp, err := client.Do(req)
	if err != nil {
		return models.EcoindexSearchResults{}, err
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return models.EcoindexSearchResults{}, err
	}

	var ecoindexes models.Ecoindexes
	err = json.Unmarshal(b, &ecoindexes)
	if err != nil {
		return models.EcoindexSearchResults{}, err
	}

	if ecoindexes.Total > 0 {
		return convertApIResult(ecoindexes.Items, host, path), nil
	}

	return models.EcoindexSearchResults{}, nil
}

func convertApIResult(ecoindexes []models.Ecoindex, host string, path string) models.EcoindexSearchResults {
	var exactResults, hostResults []models.Ecoindex

	for _, ecoindex := range ecoindexes {
		ecoindexResultUrl, err := url.Parse(ecoindex.Url)
		if err != nil {
			panic(err)
		}

		ecoindexUrl := ecoindexResultUrl.Host + ecoindexResultUrl.Path
		ecoindex.Color = GetColor(ecoindex.Grade)
		if ecoindexUrl == host+path || ecoindexUrl == strings.TrimSuffix(host+path, "/") || ecoindexUrl == host+path+"/" {
			exactResults = append(exactResults, ecoindex)
		} else {
			hostResults = append(hostResults, ecoindex)
		}
	}

	searchResults := models.EcoindexSearchResults{
		Count:       len(ecoindexes),
		HostResults: hostResults,
	}

	if len(exactResults) > 0 {
		searchResults.LatestResult = exactResults[0]
	}

	if len(exactResults) > 1 {
		searchResults.OlderResults = exactResults[1:]
	}

	return searchResults
}
