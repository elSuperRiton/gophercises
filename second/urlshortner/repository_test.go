package urlshortner

import (
	"net/http"
	"reflect"
	"testing"
)

type testRepo struct{}

func (tr testRepo) Handler(fallbackHandler http.Handler) (http.HandlerFunc, error) {
	return func(w http.ResponseWriter, r *http.Request) {
		fallbackHandler.ServeHTTP(w, r)
	}, nil
}

func Test_SetRepository(t *testing.T) {
	t.Run("Testing set repository", func(t *testing.T) {
		SetRepository(testRepo{})

		if reflect.TypeOf(repo).String() != reflect.TypeOf(testRepo{}).String() {
			t.Errorf("Wanted repo to be set with %v, got %v", reflect.TypeOf(testRepo{}).String(), reflect.TypeOf(repo).String())
		}
	})
}

func Test_Handler(t *testing.T) {

	SetRepository(nil)
	handlerFunc := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})

	t.Run("Testing Handler with repository initialized", func(t *testing.T) {

		defer func() {
			if err := recover(); err == nil {
				t.Errorf("Wanted err not to be nil")
			}
		}()

		repo.Handler(handlerFunc)
	})

	t.Run("Testing Handler with non initialized repository", func(t *testing.T) {

		SetRepository(testRepo{})
		handler, err := Handler(handlerFunc)

		if err != nil {
			t.Errorf("Wanted err to be nil, got %v", err)
		}

		if handler == nil {
			t.Error("Wanted handler not to be nil")
		}
	})
}
