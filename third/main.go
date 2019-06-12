package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/elSuperRiton/gophercices/third/story"
)

var (
	filePath        string
	storyRepository story.Repository
)

func init() {
	flag.StringVar(&filePath, "file", "./gopher.json", "absolute path to story file")
	flag.Parse()

	initStoryRepository()
}

func main() {

	srv := http.Server{
		Addr:    ":8080",
		Handler: newRouter(),
	}

	log.Fatal(srv.ListenAndServe())
}

func initStoryRepository() {
	var err error
	storyRepository, err = story.NewRepositoryFromFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
}

func newRouter() *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/", story.Handler(storyRepository))
	return mux
}
