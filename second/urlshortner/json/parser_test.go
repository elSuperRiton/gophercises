package json

import (
	"io/ioutil"
	"path/filepath"
	"testing"
)

var testRepository Repository

func helperLoadData(t *testing.T, name string) (bytes []byte) {
	path := filepath.Join("testdata", name)
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}

	return bytes
}

func setupStubTest(t *testing.T, dataType string) func(t *testing.T) {
	t.Log("setup stub test")
	testRepository = Repository{
		data: helperLoadData(t, dataType),
	}
	return func(t *testing.T) {
		t.Log("teardown test")
		testRepository = Repository{}
	}
}

func Test_Parse(t *testing.T) {

	t.Run("Testing proper JSON", func(t *testing.T) {

		teardownStubTest := setupStubTest(t, "good.json")
		defer teardownStubTest(t)

		parsedData, err := parse(testRepository.data)
		if err != nil {
			t.Errorf("Wanted error to be nil, go %v", err)
		}

		if len(parsedData) != 2 {
			t.Errorf("Wanted lenght of parseYAML to be 2, got %v", len(parsedData))
		}
	})

	t.Run("Testing malformed JSON", func(t *testing.T) {

		teardownStubTest := setupStubTest(t, "bad.json")
		defer teardownStubTest(t)

		_, err := parse(testRepository.data)
		if err == nil {
			t.Errorf("Wanted err not be to nil")
		}
	})

}
