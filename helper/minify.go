package helper

import (
	"log"

	"github.com/cnumr/ecoindex-bff/config"
)

func MinifyString(mediaType string, input string) string {
	minified, err := config.MINIFIER.String(mediaType, input)
	if err != nil {
		log.Default().Println(err)
		return input
	}

	return minified
}
