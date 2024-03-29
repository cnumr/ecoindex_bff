package config

import (
	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/svg"
)

var MINIFIER *minify.M

func GetMinifier() *minify.M {
	m := minify.New()
	m.AddFunc("image/svg+xml", svg.Minify)

	return m
}
