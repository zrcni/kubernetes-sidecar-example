package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"text/template"
)

var imgDir = os.Getenv("IMAGE_DIR")

type PageData struct {
	Images []Image
	Title  string
}

type Image struct {
	Name   string
	Source string
}

func makeImageURL(name string) string {
	return fmt.Sprintf("/images/%s", name)
}

func imagesHandler(w http.ResponseWriter, r *http.Request) {
	files, err := ioutil.ReadDir(imgDir)
	if err != nil {
		log.Fatal(err)
	}
	var images []Image
	for _, f := range files {
		image := Image{
			Name:   f.Name(),
			Source: makeImageURL(f.Name()),
		}
		images = append(images, image)
	}

	tmpl, err := template.ParseFiles("/app/layout.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err.Error())
		return
	}

	data := PageData{
		Images: images,
		Title:  "app",
	}
	tmpl.Execute(w, data)
}

func main() {
	fs := http.FileServer(http.Dir(imgDir))
	http.Handle("/images/", http.StripPrefix("/images/", fs))

	http.HandleFunc("/", imagesHandler)
	log.Fatal(http.ListenAndServe(":3000", nil))

	<-exitOnSignal(syscall.SIGINT, syscall.SIGTERM)
}

func exitOnSignal(sigs ...os.Signal) chan bool {
	sigsChan := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	signal.Notify(sigsChan, sigs...)

	go func() {
		<-sigsChan
		done <- true
	}()

	return done
}
