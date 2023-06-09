package handlers

import (
	"log"
	"net/http"
	"os"
	"sort"
	"strings"
	"time"
)

func Videos(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")

	log.Printf("VideosIndex %s request %s", r.URL, parts[2])

	if parts[2] != "" {
		renderPlayer(w, parts[2])
	} else {
		renderIndex(w)
	}
}

func renderPlayer(w http.ResponseWriter, filename string) {
	vf := VideoFile{
		Name: filename,
	}

	err := PlayerTemplate.Execute(w, vf)

	if err != nil {
		log.Printf("error rendering player: %s", err)
	}
}

func renderIndex(w http.ResponseWriter) {
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
	Name    string
	ModTime time.Time
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
			info, _ := file.Info()
			vfs = append(vfs, VideoFile{
				Name:    file.Name(),
				ModTime: info.ModTime(),
			})
		}
	}

	sort.Slice(vfs, func(i int, j int) bool {
		return vfs[i].ModTime.After(vfs[j].ModTime)
	})

	return vfs, nil
}
