package handlers

import (
	"html/template"
	"log"
)

type TemplateMap map[string]*template.Template

var Templates TemplateMap

var PlayerTemplate *template.Template

func init() {
	log.Printf("Init Templates")
	Templates = TemplateMap{
		"videos": template.Must(template.ParseFiles("templates/videos.html")),
	}
	PlayerTemplate = template.Must(template.ParseFiles("templates/player.html"))
}
