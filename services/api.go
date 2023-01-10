package services

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/vvatelot/ecoindex-microfront/config"
	"github.com/vvatelot/ecoindex-microfront/models"
)

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
	resultCount := len(ecoindexes)

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
		Count:       resultCount,
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
