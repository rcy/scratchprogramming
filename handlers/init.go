package handlers

import (
	_ "embed"
	"html/template"
)

//go:embed "templates/player.html"
var playerTemplateContent string
var PlayerTemplate = template.Must(template.New("").Parse(playerTemplateContent))

//go:embed "templates/videos.html"
var videosTemplateContent string
var VideosTemplate = template.Must(template.New("").Parse(videosTemplateContent))
