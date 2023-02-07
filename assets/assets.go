package assets

import "embed"

//go:embed template/*
var TemplateFs embed.FS

//go:embed js/*
var JsFs embed.FS
