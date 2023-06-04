package handlers

import (
	"log"
	"net/http"
	"os"
)

func Videos(w http.ResponseWriter, r *http.Request) {
	log.Printf("Videos %s request", r.URL)

	// collect list of video filenames from assets/videos

	vfs, err := getVideos()
	if err != nil {
		w.WriteHeader(500)
	}

	err = Templates["videos"].Execute(w, struct {
		VideoFiles []VideoFile
	}{
		VideoFiles: vfs,
	})

	if err != nil {
		w.WriteHeader(500)
	}
}

type VideoFile struct {
	Name string
}

func getVideos() ([]VideoFile, error) {
	dir := "/home/rcy/src/scratchprogramming/assets/videos"

	vfs := []VideoFile{}

	files, err := os.ReadDir(dir)
	if err != nil {
		return []VideoFile{}, err
	}

	for _, file := range files {
		if !file.IsDir() {
			vfs = append(vfs, VideoFile{Name: file.Name()})
		}
	}

	return vfs, nil
}
