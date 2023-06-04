package handlers

import (
	"html/template"
	"log"
)

type TemplateMap map[string]*template.Template

var Templates TemplateMap

func init() {
	log.Printf("Init Templates")
	Templates = TemplateMap{
		"home":   template.Must(template.ParseFiles("templates/home.html")),
		"videos": template.Must(template.ParseFiles("templates/videos.html")),
	}
}
