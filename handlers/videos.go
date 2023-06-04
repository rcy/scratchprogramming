package handlers

import (
	"log"
	"net/http"
)

func Videos(w http.ResponseWriter, r *http.Request) {
	log.Printf("Videos %s request", r.URL)

	// collect list of video filenames from assets/videos

	Templates["videos"].Execute(w, nil)
}
