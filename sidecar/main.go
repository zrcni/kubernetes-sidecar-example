package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path"
	"strings"
	"syscall"
	"time"
)

var imgDir = os.Getenv("IMAGE_DIR")

func clearImageDir() {
	dir, err := ioutil.ReadDir(imgDir)
	if err != nil {
		log.Println(err)
		return
	}
	for _, file := range dir {
		os.RemoveAll(path.Join([]string{imgDir, file.Name()}...))
	}
}

func fetchRandomImage() (io.Reader, string, error) {
	response, err := http.Get("https://picsum.photos/500/500")
	if err != nil {
		return nil, "", err
	}
	path := response.Request.URL.Path
	splitPath := strings.Split(response.Request.URL.Path, "/")
	if len(splitPath) < 5 {
		return nil, "", fmt.Errorf("unexpected url path: %s", path)
	}
	id := splitPath[2]
	return response.Body, id, nil
}

func saveImage(r io.Reader, name string) error {
	file, err := os.Create(fmt.Sprintf("%s/%s.jpg", imgDir, name))
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, r)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	go func() {
		for {
			clearImageDir()
			log.Println("Removed images")
			imageIds := map[string]string{}

			count := 0
			for count < 5 {
				body, id, err := fetchRandomImage()
				if err != nil {
					log.Println(err)
					continue
				}

				if imageIds[id] != "" {
					log.Printf("Image %s already exists. Fetching another one.", id)
					continue
				}

				err = saveImage(body, id)
				if err != nil {
					log.Println(err)
					continue
				}
				count++
			}
			log.Println("Saved images")

			time.Sleep(60 * time.Second)
		}
	}()

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
