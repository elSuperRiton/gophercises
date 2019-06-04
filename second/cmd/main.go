package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/elSuperRiton/gophercices/second/urlshortner/json"
	"github.com/elSuperRiton/gophercices/second/urlshortner/yaml"

	"github.com/elSuperRiton/gophercices/second/urlshortner"
)

const (
	serverPort = ":8080"
)

func main() {
	var parserType string
	flag.StringVar(&parserType, "parser", "yaml", "the needed parser ( \"yaml\" || \"json\" || \"bolt\" )")
	flag.Parse()

	mux, err := muxWithHandler(parserType)
	if err != nil {
		panic(err)
	}

	srv := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("error starting server on port %v : %v", serverPort, err)
	}
}

// muxWithHandler returns a SeveMux with a proper handler attached
// to the root route
func muxWithHandler(parser string) (*http.ServeMux, error) {

	switch parser {
	case "yaml":
		urlshortner.SetRepository(yaml.NewRepository(testProperlyformedYAML))
	case "json":
		urlshortner.SetRepository(json.NewRepository(testMalformedJSON))
	case "bolt":
	default:
		return nil, fmt.Errorf("Please provide a valid parser value")
	}

	hdl, err := urlshortner.Handler(http.HandlerFunc(func(w http.ResponseWriter, t *http.Request) {
		w.Write(defaultResponse)
	}))
	if err != nil {
		return nil, err
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", hdl)

	return mux, nil
}
