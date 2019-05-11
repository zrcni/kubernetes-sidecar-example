package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"text/template"
	"time"
)

var imgDir = os.Getenv("IMAGE_DIR")

type PageData struct {
	Images          []Image
	Title           string
	ImagesUpdatedAt string
}

type Image struct {
	Name   string
	Source string
}

func getExecPath() string {
	ex, err := os.Executable()
	if err != nil {
		log.Println(err)
		return "/"
	}
	exPath := filepath.Dir(ex)
	return exPath
}

func makeImageURL(name string) string {
	return fmt.Sprintf("/images/%s", name)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	files, err := ioutil.ReadDir(imgDir)
	if err != nil {
		log.Fatal(err)
	}

	var imagesUpdatedAt time.Time
	var images []Image

	for _, f := range files {
		image := Image{
			Name:   f.Name(),
			Source: makeImageURL(f.Name()),
		}
		images = append(images, image)

		fileModifiedAt := f.ModTime()
		if imagesUpdatedAt.Unix() < fileModifiedAt.Unix() {
			imagesUpdatedAt = fileModifiedAt
		}
	}

	tmpl, err := template.ParseFiles("./templates/index.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err.Error())
		return
	}

	data := PageData{
		Images:          images,
		ImagesUpdatedAt: imagesUpdatedAt.Format("2/1/2006 15:04:05 MST"),
		Title:           "kubernetes-sidecar-example",
	}
	tmpl.Execute(w, data)
}

func main() {
	fs := http.FileServer(http.Dir(imgDir))
	http.Handle("/images/", http.StripPrefix("/images/", fs))

	http.HandleFunc("/", indexHandler)

	log.Println("Listening on port 3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
