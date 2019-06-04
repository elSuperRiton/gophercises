package json

import (
	"net/http"

	"github.com/elSuperRiton/gophercices/second/urlshortner/utils"
)

// Handler parses the provided JSON and then return an http.HandlerFunc
// that will attempt to map any paths to their corresponding URL.
// If the path is not provided in the JSON, then the fallback http.Handler will be called
//
// JSON is expected to be in the following format :
//
// [{
// 	"path": "/some-path",
// 	"url": "https://the-url-to-redirect-to.com/well-done",
// },
// {
// 	"path": "/some-path",
// 	"url": "https://the-url-to-redirect-to.com/well-done",
// }]
//
// The only errors that can be returned all relate to having an invalid YAML format
//
// See MapHandler to create a similar http.HandlerFunc via a mapping of paths to urls.
func (r Repository) Handler(fallback http.Handler) (http.HandlerFunc, error) {
	u, err := parse(r.data)
	if err != nil {
		return nil, err
	}

	return utils.MapHandler(utils.BuildMap(u), fallback), nil
}
