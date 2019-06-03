package utils

import (
	"errors"
	"testing"
)

func TestPanicIfErr(t *testing.T) {

	testMessage := "message should be passed down"
	testError := errors.New("Error should cause a panic")

	t.Run("Testing the PanicIfErr function with non nil error", func(t *testing.T) {

		defer func() {
			if r := recover(); r == nil {
				t.Error("Wanted panic called. got nil")
			}
		}()

		PanicIfErr(testError, testMessage)
	})

	t.Run("Testing the PanicIfErr function with nil error", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				t.Error("Wanted nothing, got a panic")
			}
		}()

		PanicIfErr(nil, testMessage)
	})

}
