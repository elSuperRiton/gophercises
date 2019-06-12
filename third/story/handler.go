package story

import (
	"net/http"
)

// Handler returns an http Handler that extracts the potentially
// selected chapter from the request and executes the template
// in order to render it
func Handler(repo Repository) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		repo.ChapterFromRequest(r)

		layoutTpl.Execute(w, struct {
			ChapterTitle ChapterTitle
			Chapter      Chapter
		}{
			ChapterTitle: repo.Current(),
			Chapter:      repo.CurrentChapter(),
		})
	})
}
