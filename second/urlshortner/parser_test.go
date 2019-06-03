package urlshortner

import (
	"strconv"
	"testing"
)

var properlyformedYAML = []byte(`
- path: /urlshort
  url: https://github.com/gophercises/urlshort
- path: /urlshort-final
  url: https://github.com/gophercises/urlshort/tree/solution
`)

var malformedYAML = []byte(`
	- path: /urlshort
		url: https://github.com/gophercises/urlshort
	- path: /urlshort-final
		url: https://github.com/gophercises/urlshort/tree/solution
`)

func Test_parseYAML(t *testing.T) {

	t.Run("Testing proper YAML", func(t *testing.T) {
		parsedYAML, err := parseYAML(properlyformedYAML)
		if err != nil {
			t.Errorf("Wanted error to be nil, go %v", err)
		}

		if len(parsedYAML) != 2 {
			t.Errorf("Wanted lenght of parseYAML to be 2, got %v", len(parsedYAML))
		}
	})

	t.Run("Testing malformed YAML", func(t *testing.T) {
		_, err := parseYAML(malformedYAML)
		if err == nil {
			t.Errorf("Wanted err not be to nil")
		}
	})

}

func Test_buildMap(t *testing.T) {
	t.Run("Testing buildMap", func(t *testing.T) {

		var u urls
		for index := 0; index < 10; index++ {
			url := make(map[string]string)
			url["path"] = "/short-path-" + strconv.Itoa(index)
			url["url"] = "http://long-path-" + strconv.Itoa(index) + ".com"
			u = append(u, url)
		}

		mappedURLS := buildMap(u)
		for _, url := range u {
			if mappedURLS[url["path"]] != url["url"] {
				t.Errorf("Wanted %v, got %v", url["url"], mappedURLS[url["path"]])
			}
		}
	})
}
