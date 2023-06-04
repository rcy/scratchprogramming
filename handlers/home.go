package handlers

import (
	"log"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.String() != "/" {
		log.Printf("Home %s request: 404", r.URL)
		w.WriteHeader(404)
		return
	}

	log.Printf("Home %s request", r.URL)

	Templates["home"].Execute(w, nil)
}
