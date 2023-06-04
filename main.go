package main

import (
	"html/template"
	"log"
	"net/http"
)

type TemplateMap map[string]*template.Template

var t TemplateMap

func main() {
	http.HandleFunc("/", getHome)
	http.HandleFunc("/videos", getVideos)
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	t = initTemplates()

	log.Print("listening on 3333...")
	err := http.ListenAndServe(":3333", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func initTemplates() TemplateMap {
	return TemplateMap{
		"home":   template.Must(template.ParseFiles("templates/home.html")),
		"videos": template.Must(template.ParseFiles("templates/videos.html")),
	}
}

func getHome(w http.ResponseWriter, r *http.Request) {
	if r.URL.String() != "/" {
		log.Printf("getHome %s request: 404", r.URL)
		w.WriteHeader(404)
		return
	}

	log.Printf("getHome %s request", r.URL)

	t["home"].Execute(w, nil)
}

func getVideos(w http.ResponseWriter, r *http.Request) {
	log.Printf("getVideos %s request", r.URL)
	t["videos"].Execute(w, nil)
}
