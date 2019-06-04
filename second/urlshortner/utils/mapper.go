package utils

import "net/http"

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

// BuildMap takes in a []map[string]string and maps
// the short path values to their corresponding full path
func BuildMap(u []map[string]string) map[string]string {
	mappedURLS := make(map[string]string)
	for _, url := range u {
		mappedURLS[url["path"]] = url["url"]
	}

	return mappedURLS
}
