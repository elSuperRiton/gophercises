package utils

import (
	"net/http"
	"net/http/httptest"
	"strconv"
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

func Test_BuildMap(t *testing.T) {
	t.Run("Testing buildMap", func(t *testing.T) {

		var u []map[string]string
		for index := 0; index < 10; index++ {
			url := make(map[string]string)
			url["path"] = "/short-path-" + strconv.Itoa(index)
			url["url"] = "http://long-path-" + strconv.Itoa(index) + ".com"
			u = append(u, url)
		}

		mappedURLS := BuildMap(u)
		for _, url := range u {
			if mappedURLS[url["path"]] != url["url"] {
				t.Errorf("Wanted %v, got %v", url["url"], mappedURLS[url["path"]])
			}
		}
	})
}
