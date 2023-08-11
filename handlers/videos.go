package handlers

import (
	"log"
	"net/http"
	"net/url"
	"os"
	"sort"
	"strings"
	"time"
)

func Videos(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")

	log.Printf("VideosIndex %s request %s", r.URL, parts)

	if len(parts) == 4 {
		filename, err := url.QueryUnescape(parts[3])
		if err != nil {
			http.Error(w, err.Error(), 500)
		}
		renderPlayer(w, parts[2], filename)
	} else {
		renderIndex(w, parts[2])
	}
}

func renderPlayer(w http.ResponseWriter, subdir string, filename string) {
	vf := VideoFile{
		SubDir: subdir,
		Name:   filename,
	}

	subdirs, err := getSubDirs()
	if err != nil {
		w.WriteHeader(500)
	}

	err = PlayerTemplate.Execute(w, struct {
		SubDirs   []string
		VideoFile VideoFile
	}{
		SubDirs:   subdirs,
		VideoFile: vf,
	})

	if err != nil {
		log.Printf("error rendering player: %s", err)
	}
}

func renderIndex(w http.ResponseWriter, subdir string) {
	// collect list of video filenames from assets/videos

	vfs, err := getVideos(subdir)
	if err != nil {
		w.WriteHeader(500)
	}

	subdirs, err := getSubDirs()
	if err != nil {
		w.WriteHeader(500)
	}

	err = VideosTemplate.Execute(w, struct {
		SubDir     string
		VideoFiles []VideoFile
		SubDirs    []string
	}{
		SubDir:     subdir,
		VideoFiles: vfs,
		SubDirs:    subdirs,
	})

	if err != nil {
		w.WriteHeader(500)
	}
}

type VideoFile struct {
	SubDir  string
	Name    string
	ModTime time.Time
}

func (vf VideoFile) EscapedName() string {
	return url.QueryEscape(vf.Name)
}

func getVideos(subdir string) ([]VideoFile, error) {
	dir := os.Getenv("ASSET_DIRECTORY") + "/videos/" + subdir

	log.Printf("getVideos dir=%s", dir)

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

func getSubDirs() ([]string, error) {
	dir := os.Getenv("ASSET_DIRECTORY") + "/videos/"

	subdirs := []string{}

	files, err := os.ReadDir(dir)
	if err != nil {
		return []string{}, err
	}

	for _, file := range files {
		if file.IsDir() {
			subdirs = append(subdirs, file.Name())
		}
	}

	sort.Slice(subdirs, func(i int, j int) bool {
		return subdirs[i] > subdirs[j]
	})

	return subdirs, nil
}
