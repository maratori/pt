package pt_test

import (
	"testing"
	"time"

	"github.com/maratori/pt"
)

func TestParallel(t *testing.T) {
	t.Parallel()
	t.Run("should panic on nil T", func(t *testing.T) {
		t.Parallel()
		defer assertPanic(t, "argument t *testing.T can not be nil")
		pt.Parallel(nil)
	})
	t.Run("should not panic without tests", func(t *testing.T) {
		t.Parallel()
		pt.Parallel(&testing.T{})
	})
	t.Run("should not panic with 1 test", func(t *testing.T) {
		t.Parallel()
		pt.Parallel(t, testing.InternalTest{F: func(t *testing.T) {}})
	})
	t.Run("should not panic with 2 tests", func(t *testing.T) {
		t.Parallel()
		pt.Parallel(t, testing.InternalTest{F: func(t *testing.T) {}}, testing.InternalTest{F: func(t *testing.T) {}})
	})
	t.Run("should run 1 test", func(t *testing.T) {
		t.Parallel()
		called := false
		t.Run("internal", func(t *testing.T) {
			pt.Parallel(t, testing.InternalTest{F: func(t *testing.T) {
				if called {
					t.Error("test is called twice")
				}
				called = true
			}})
		})
		if !called {
			t.Error("test is not called")
		}
	})
	t.Run("should run 2 tests parallel", func(t *testing.T) {
		singleTestTime := 50 * time.Millisecond
		timeout := singleTestTime + 20*time.Millisecond
		called1 := false
		called2 := false
		start := time.Now()
		t.Run("internal", func(t *testing.T) {
			pt.Parallel(t,
				testing.InternalTest{F: func(t *testing.T) {
					if called1 {
						t.Error("test1 is called twice")
					}
					called1 = true
					time.Sleep(singleTestTime)
				}},
				testing.InternalTest{F: func(t *testing.T) {
					if called2 {
						t.Error("test2 is called twice")
					}
					called2 = true
					time.Sleep(singleTestTime)
				}},
			)
		})
		elapsed := time.Since(start)
		if elapsed < singleTestTime {
			t.Errorf("tests execution time %s not exceeded %s", elapsed, singleTestTime)
		}
		if elapsed > timeout {
			t.Errorf("tests execution time %s exceeded %s", elapsed, timeout)
		}
		if !called1 {
			t.Error("test1 is not called")
		}
		if !called2 {
			t.Error("test2 is not called")
		}
	})
	t.Run("should run 100 tests parallel", func(t *testing.T) {
		singleTestTime := 50 * time.Millisecond
		timeout := singleTestTime * 100 / 2
		tests := make([]testing.InternalTest, 100)
		for i := range tests {
			tests[i] = testing.InternalTest{F: func(t *testing.T) {
				time.Sleep(singleTestTime)
			}}
		}
		start := time.Now()
		t.Run("internal", func(t *testing.T) {
			pt.Parallel(t, tests...)
		})
		elapsed := time.Since(start)
		if elapsed < singleTestTime {
			t.Errorf("tests execution time %s not exceeded %s", elapsed, singleTestTime)
		}
		if elapsed > timeout {
			t.Errorf("tests execution time %s exceeded %s", elapsed, timeout)
		}
	})
}

func TestGroup(t *testing.T) {
	t.Parallel()
	t.Run("should return right name", func(t *testing.T) {
		t.Parallel()
		internalTest := pt.Group("abc")
		if internalTest.Name != "abc" {
			t.Error("name is wrong")
		}
	})
	t.Run("should not panic without tests", func(t *testing.T) {
		t.Parallel()
		internalTest := pt.Group("")
		internalTest.F(&testing.T{})
	})
	t.Run("should run 1 test", func(t *testing.T) {
		t.Parallel()
		called := false
		internalTest := pt.Group("", testing.InternalTest{F: func(t *testing.T) {
			if called {
				t.Error("test is called twice")
			}
			called = true
		}})
		if called {
			t.Error("test is called too early")
		}
		t.Run("internal", func(t *testing.T) {
			internalTest.F(t)
		})
		if !called {
			t.Error("test is not called")
		}
	})
	t.Run("should run 2 tests parallel", func(t *testing.T) {
		singleTestTime := 50 * time.Millisecond
		timeout := singleTestTime + 20*time.Millisecond
		called1 := false
		called2 := false
		internalTest := pt.Group("",
			testing.InternalTest{F: func(t *testing.T) {
				if called1 {
					t.Error("test1 is called twice")
				}
				called1 = true
				time.Sleep(singleTestTime)
			}},
			testing.InternalTest{F: func(t *testing.T) {
				if called2 {
					t.Error("test2 is called twice")
				}
				called2 = true
				time.Sleep(singleTestTime)
			}},
		)
		start := time.Now()
		t.Run("internal", func(t *testing.T) {
			internalTest.F(t)
		})
		elapsed := time.Since(start)
		if elapsed < singleTestTime {
			t.Errorf("tests execution time %s not exceeded %s", elapsed, singleTestTime)
		}
		if elapsed > timeout {
			t.Errorf("tests execution time %s exceeded %s", elapsed, timeout)
		}
		if !called1 {
			t.Error("test1 is not called")
		}
		if !called2 {
			t.Error("test2 is not called")
		}
	})

}

func TestTest(t *testing.T) {
	t.Parallel()
	t.Run("should panic on nil", func(t *testing.T) {
		t.Parallel()
		defer assertPanic(t, "argument test func(t *testing.T) can not be nil")
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
		testFunc := func(_ *testing.T) {
			if called {
				t.Error("test is called twice")
			}
			called = true
		}
		internalTest := pt.Test("", testFunc)
		if called {
			t.Error("test is called too early")
		}
		internalTest.F(nil)
		if !called {
			t.Error("test is not called")
		}
	})
}

func assertPanic(t *testing.T, expected string) {
	if value := recover(); value == nil {
		t.Error("no panic")
	} else if value != expected {
		t.Errorf("unexpected panic value: %v", value)
	}
}
