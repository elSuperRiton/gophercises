package yaml

import (
	"reflect"
	"testing"
)

func Test_NewRepository(t *testing.T) {
	t.Run("Testing NewRepository function", func(t *testing.T) {
		testData := []byte("Test data")

		r := NewRepository(testData)

		if r.data == nil {
			t.Errorf("Wanted r not to be nil")
		}

		if !reflect.DeepEqual(r.data, testData) {
			t.Errorf("Wanted r.data to be %v, got %v", testData, r.data)
		}
	})

}
