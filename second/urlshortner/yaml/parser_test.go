package yaml

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
	testRepository = Repository{
		data: helperLoadData(t, dataType),
	}
	return func(t *testing.T) {
		testRepository = Repository{}
	}
}
func Test_Parse(t *testing.T) {

	t.Run("Testing proper YAML", func(t *testing.T) {

		tearDownStubTest := setupStubTest(t, "good.yaml")
		defer tearDownStubTest(t)

		parsedYAML, err := parse(testRepository.data)
		if err != nil {
			t.Errorf("Wanted error to be nil, go %v", err)
		}

		if len(parsedYAML) != 2 {
			t.Errorf("Wanted lenght of parseYAML to be 2, got %v", len(parsedYAML))
		}
	})

	t.Run("Testing malformed YAML", func(t *testing.T) {

		tearDownStubTest := setupStubTest(t, "bad.yaml")
		defer tearDownStubTest(t)

		_, err := parse(testRepository.data)
		if err == nil {
			t.Errorf("Wanted err not be to nil")
		}
	})

}
