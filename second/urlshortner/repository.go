package urlshortner

import (
	"net/http"
)

// Repository is an interface that needs to be satisfied by data Handlers
type Repository interface {
	Handler(http.Handler) (http.HandlerFunc, error)
}

var repo Repository

// SetRepository initializes the urlshortner with a
// passed in Repository interface
func SetRepository(r Repository) {
	repo = r
}

// Handler calls the Repository implementation Handler method
func Handler(fallback http.Handler) (http.HandlerFunc, error) {
	return repo.Handler(fallback)
}
