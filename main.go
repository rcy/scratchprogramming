package main

import (
	"log"
	"net/http"
	"scrapro/handlers"
)

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/videos/", handlers.Videos)
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	log.Print("listening on 3333...")
	err := http.ListenAndServe(":3333", nil)
	if err != nil {
		log.Fatal(err)
	}
}
