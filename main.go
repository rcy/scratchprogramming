package main

import (
	"log"
	"net/http"
	"net/url"
	"os"
	"scrapro/handlers"
)

func unescape(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		p, err := url.QueryUnescape(r.URL.Path)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		rp, err := url.QueryUnescape(r.URL.RawPath)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		r2 := new(http.Request)
		*r2 = *r
		r2.URL = new(url.URL)
		*r2.URL = *r.URL
		r2.URL.Path = p
		r2.URL.RawPath = rp
		h.ServeHTTP(w, r2)
	})
}

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/videos/", handlers.Videos)

	assetDirectory := http.Dir(os.Getenv("ASSET_DIRECTORY"))
	log.Printf("assetDirectory: %s", assetDirectory)
	http.Handle("/assets/", http.StripPrefix("/assets/", unescape(http.FileServer(assetDirectory))))

	log.Print("listening on 3333...")
	err := http.ListenAndServe(":3333", nil)
	if err != nil {
		log.Fatal(err)
	}
}
