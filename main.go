package main

import (
	"log"
	"net/http"
	"os"
	"scrapro/handlers"
)

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/videos/", handlers.Videos)

	assetDirectory := http.Dir(os.Getenv("ASSET_DIRECTORY"))
	log.Printf("assetDirectory: %s", assetDirectory)
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(assetDirectory)))

	log.Print("listening on 3333...")
	err := http.ListenAndServe(":3333", nil)
	if err != nil {
		log.Fatal(err)
	}
}
