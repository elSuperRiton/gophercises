package boltdb

import (
	"net/http"
)

// Handler attempts to map any paths to their corresponding URL.
// If the path doesn't exists within bucket, then the fallback http.Handler will be called
func (repo Repository) Handler(fallback http.Handler) (http.HandlerFunc, error) {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if url := repo.matchShortURL(r.URL.Path); len(url) > 0 {
			http.Redirect(w, r, string(url), http.StatusFound)
		}

		fallback.ServeHTTP(w, r)
	}), nil
}
