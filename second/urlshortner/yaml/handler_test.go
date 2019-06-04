package yaml

import (
	"net/http"
	"testing"
)

func Test_Handler(t *testing.T) {

	fallbackHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Success"))
	})

	t.Run("Test YAMLHandler with proper YAML", func(t *testing.T) {

		tearDownStubTest := setupStubTest(t, "good.yaml")
		defer tearDownStubTest(t)

		handler, err := testRepository.Handler(fallbackHandler)
		if err != nil {
			t.Errorf("Wanted err to be nil, got %v", err)
		}

		if handler == nil {
			t.Errorf("Wanted an http.Handler to be returned")
		}
	})

	t.Run("Test YAMLHandler with malformed YAML", func(t *testing.T) {

		tearDownStubTest := setupStubTest(t, "bad.yaml")
		defer tearDownStubTest(t)

		if _, err := testRepository.Handler(fallbackHandler); err == nil {
			t.Error("Wanted error from YAMLHandler not to be nil")
		}
	})
}
