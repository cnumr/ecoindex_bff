package services

import (
	"io"
	"net/http"

	"github.com/cnumr/ecoindex-bff/config"
	"github.com/cnumr/ecoindex-bff/helper"
)

func GetColor(grade string) string {
	colors := map[string]string{
		"A": "#349A47",
		"B": "#51B84B",
		"C": "#CADB2A",
		"D": "#F6EB15",
		"E": "#FECD06",
		"F": "#F99839",
		"G": "#ED2124",
	}

	if color, ok := colors[grade]; ok {
		return color
	}

	return "lightgrey"
}

func GetBadgeSvg(grade string, theme string) string {
	url := config.ENV.CDNUrl + "@" + config.ENV.BadgeVersion + "/assets/svg/" + theme + "/" + grade + ".svg"
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	badgeSvg, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	return helper.MinifyString("image/svg+xml", string(badgeSvg))
}
