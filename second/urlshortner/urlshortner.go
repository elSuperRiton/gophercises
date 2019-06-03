package urlshortner

import (
	"net/http"
)

// MapHandler returns an http.HandlerFunc that will attempt to
// map any paths to their corresponding URL
// If the path is not provided in the map, then the fallback http.Handler will
// be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if url, ok := pathsToUrls[r.URL.Path]; ok {
			http.Redirect(w, r, url, http.StatusFound)
		}

		fallback.ServeHTTP(w, r)
	}
}

// YAMLHandler parses the provided YAML and then return an http.HandlerFunc
// that will attempt to map any paths to their corresponding URL.
// If the path is not provided in the YAML, then the fallback http.Handler will be called
//
// YAML is expected to be in the following format :
//
// 		- path: /some-path
//			url: https://the-url-to-redirect-to.com/well-done
//
// The only errors that can be returned all relate to having an invalid YAML format
//
// See MapHandler to create a similar http.HandlerFunc via a mapping of paths to urls.
func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	u, err := parseYAML(yml)
	if err != nil {
		return nil, err
	}

	return MapHandler(buildMap(u), fallback), nil
}

// JSONHandler parses the provided JSON and then return an http.HandlerFunc
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
func JSONHandler(json []byte, fallback http.Handler) (http.HandlerFunc, error) {
	u, err := parseJSON(json)
	if err != nil {
		return nil, err
	}

	return MapHandler(buildMap(u), fallback), nil
}
