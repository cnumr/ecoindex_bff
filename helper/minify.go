package helper

import "github.com/cnumr/ecoindex-bff/config"

func MinifyString(mediaType string, input string) string {
	minified, err := config.MINIFIER.String(mediaType, input)
	if err != nil {
		panic(err)
	}

	return minified
}
