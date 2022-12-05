package main

import (
	"embed"
	"io/fs"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	"github.com/vvatelot/ecoindex-microfront/config"
	"github.com/vvatelot/ecoindex-microfront/handler"
)

var ENV *config.Environment = config.GetEnvironment()

//go:embed views/*
var embedDirViews embed.FS

func main() {
	config.ENV = ENV

	viewsFileSystem, err := fs.Sub(embedDirViews, "views")
	if err != nil {
		panic(err)
	}
	app := fiber.New(fiber.Config{
		Views: html.NewFileSystem(http.FS(viewsFileSystem), ".html"),
	})

	app.Get("/", handler.GetEcoindexBadge)

	app.Listen(":" + config.ENV.AppPort)
}
