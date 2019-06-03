package urlshortner

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func redirectCaptureClient() *http.Client {
	return &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
}

func Test_MapHandler(t *testing.T) {

	shortURIMap := make(map[string]string)
	shortURIMap["/short-url"] = "https://google.com"
	mapHandler := http.HandlerFunc(MapHandler(shortURIMap, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})))
	testServer := httptest.NewServer(mapHandler)

	t.Run("Test MapHandler with existing short url", func(t *testing.T) {

		// declare an http client with a CheckRedirect function returning a http.ErrUseLastResponse
		// in order to capture the redirection instead of the actual redirect response
		httpClient := redirectCaptureClient()

		// test request
		res, err := httpClient.Get(testServer.URL + "/short-url")
		if err != nil {
			t.Fatal(err)
		}

		if res.StatusCode != http.StatusFound {
			t.Errorf("Wanted status code to be %d, got %d\n", http.StatusFound, res.StatusCode)
		}

		if uri, _ := res.Location(); uri != nil && uri.String() != "https://google.com" {
			t.Errorf("Wanted location to be https://google.com, got %v\n", uri)
		}
	})

	t.Run("Test MapHandler calling fallback", func(t *testing.T) {

	})
}

func Test_YAMLHandler(t *testing.T) {

	fallbackHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Success"))
	})

	t.Run("Test YAMLHandler with malformed YAML", func(t *testing.T) {
		if _, err := YAMLHandler(malformedYAML, fallbackHandler); err == nil {
			t.Error("Wanted error from YAMLHandler not to be nil")
		}
	})

	t.Run("Test YAMLHandler with proper YAML", func(t *testing.T) {
		handler, err := YAMLHandler(properlyformedYAML, fallbackHandler)
		if err != nil {
			t.Errorf("Wanted err to be nil, got %v", err)
		}

		if handler == nil {
			t.Errorf("Wanted an http.Handler to be returned")
		}
	})
}
