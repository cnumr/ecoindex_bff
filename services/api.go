package services

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/vvatelot/ecoindex-microfront/config"
	"github.com/vvatelot/ecoindex-microfront/models"
)

func GetEcoindex(url string) (models.Ecoindex, error) {
	client := http.Client{}
	req, err := http.NewRequest(http.MethodGet, config.ENV.ApiUrl+"/v1/ecoindexes", nil)
	if err != nil {
		return models.Ecoindex{}, err
	}

	q := req.URL.Query()
	q.Add("host", url)

	req.URL.RawQuery = q.Encode()
	req.Header = http.Header{
		"x-rapidapi-key": {"3037e7e96fmsh12bedced9f019f8p1cd804jsn4967070f8bda"},
	}

	resp, err := client.Do(req)
	if err != nil {
		return models.Ecoindex{}, err
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return models.Ecoindex{}, err
	}

	var ecoindexes models.Ecoindexes
	err = json.Unmarshal(b, &ecoindexes)
	if err != nil {
		return models.Ecoindex{}, err
	}

	if ecoindexes.Total > 0 {
		return ecoindexes.Items[0], nil
	}

	return models.Ecoindex{}, nil
}
