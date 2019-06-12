package story

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	"golang.org/x/xerrors"
)

// Repository is an interface representing
// a story line
type Repository interface {
	Current() ChapterTitle
	CurrentChapter() Chapter
	ChapterFromRequest(*http.Request)
}

type repository struct {
	current ChapterTitle
	story   Story
}

// NewRepositoryFromFile returns a new repository instanciated with the
// content of a properly formated json file
func NewRepositoryFromFile(fileName string) (Repository, error) {
	r := &repository{}
	if err := r.LoadFromFile(fileName); err != nil {
		return nil, err
	}

	return r, nil
}

// Current is a helper that returns the current chapter title
// If the current chapter title is empty it then returns the default
// chapter title ( aka the initial one "intro")
func (r *repository) Current() ChapterTitle {
	if r.current == "" {
		r.current = "intro"
	}

	return r.current
}

// CurrentChapter returns the current chapter
func (r *repository) CurrentChapter() Chapter {
	return r.story[r.Current()]
}

// ChapterFromRequest checks the request for a query param.
// If the param is set it then sets the current chapter to the provided one
// else it sets the current chapter to "intro"
func (r *repository) ChapterFromRequest(request *http.Request) {
	arc := request.URL.Query().Get("arc")

	if arc != "" {
		r.current = ChapterTitle(arc)
		return
	}

	r.current = "intro"
}

// LoadFromFile takes in a file location for a story formated in JSON,
// parses that file if it exists and sets the repository story
// with the parsed content
func (r *repository) LoadFromFile(fileName string) error {

	bytes, err := ioutil.ReadFile(fileName)
	if err != nil && err != io.EOF {
		return xerrors.Errorf("error reading from json file with location %v : %v", fileName, err)
	}

	var story Story
	if err := json.Unmarshal(bytes, &story); err != nil {
		return xerrors.Errorf("error unmarshalling : %w", StoryUnmarshallErr)
	}

	r.story = story
	return nil
}
