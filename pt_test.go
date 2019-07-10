package pt_test

import (
	"testing"

	"github.com/maratori/pt"
)

func TestTest(t *testing.T) {
	t.Parallel()
	t.Run("should panic on nil", func(t *testing.T) {
		t.Parallel()
		defer func() {
			if recover() == nil {
				t.Error("func Test did not panic")
			}
		}()
		pt.Test("", nil)
	})
	t.Run("should return right name", func(t *testing.T) {
		t.Parallel()
		internalTest := pt.Test("abc", func(t *testing.T) {})
		if internalTest.Name != "abc" {
			t.Error("name is wrong")
		}
	})
	t.Run("should return right test func", func(t *testing.T) {
		t.Parallel()
		called := false
		testFunc := func(t *testing.T) {
			called = true
		}
		internalTest := pt.Test("", testFunc)
		if called {
			t.Error("test func was called")
		}
		internalTest.F(nil)
		if !called {
			t.Error("test func is wrong")
		}
	})
}
