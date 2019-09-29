package pt_test

import (
	"testing"
	"time"

	"github.com/maratori/pt"
)

func TestPackageParallel(t *testing.T) {
	t.Parallel()
	t.Run("should panic on nil T", func(t *testing.T) {
		t.Parallel()
		defer assertPanic(t, "argument t *testing.T can not be nil")
		pt.PackageParallel(nil)
	})
	t.Run("should not panic without tests", func(t *testing.T) {
		t.Parallel()
		pt.PackageParallel(t)
	})
	t.Run("should not panic with 1 test", func(t *testing.T) {
		t.Parallel()
		pt.PackageParallel(t, testing.InternalTest{F: func(*testing.T) {}})
	})
	t.Run("should not panic with 2 tests", func(t *testing.T) {
		t.Parallel()
		pt.PackageParallel(t,
			testing.InternalTest{F: func(*testing.T) {}},
			testing.InternalTest{F: func(*testing.T) {}},
		)
	})
	t.Run("should not panic if called twice", func(t *testing.T) {
		t.Parallel()
		pt.PackageParallel(t, testing.InternalTest{F: func(*testing.T) {}})
		pt.PackageParallel(t, testing.InternalTest{F: func(*testing.T) {}})
	})
	t.Run("should run 1 test", func(t *testing.T) {
		t.Parallel()
		called := false
		t.Run("internal", func(it *testing.T) {
			// internal2 is necessary because PackageParallel calls t.Parallel()
			it.Run("internal2", func(it2 *testing.T) {
				pt.PackageParallel(it2, testing.InternalTest{F: func(*testing.T) {
					if called {
						t.Error("test is called twice")
					}
					called = true
				}})
			})
		})
		if !called {
			t.Error("test is not called")
		}
	})
	t.Run("should run 2 tests", func(t *testing.T) {
		t.Parallel()
		called1 := false
		called2 := false
		t.Run("internal", func(it *testing.T) {
			// internal2 is necessary because PackageParallel calls t.Parallel()
			it.Run("internal2", func(it2 *testing.T) {
				pt.PackageParallel(it2,
					testing.InternalTest{F: func(*testing.T) {
						if called1 {
							t.Error("test1 is called twice")
						}
						called1 = true
					}},
					testing.InternalTest{F: func(*testing.T) {
						if called2 {
							t.Error("test2 is called twice")
						}
						called2 = true
					}},
				)
			})
		})
		if !called1 {
			t.Error("test1 is not called")
		}
		if !called2 {
			t.Error("test2 is not called")
		}
	})
}

func TestPackageParallel2(t *testing.T) {
	// Do not call t.Parallel() because test measures execution time
	t.Run("should run 2 tests parallel", func(t *testing.T) {
		singleTestDuration := 1 * time.Second
		expectedMaxDuration := singleTestDuration + 300*time.Millisecond
		start := time.Now()
		t.Run("internal", func(it *testing.T) {
			// internal2 is necessary because PackageParallel calls t.Parallel()
			it.Run("internal2", func(it2 *testing.T) {
				pt.PackageParallel(it2,
					testing.InternalTest{F: func(*testing.T) {
						time.Sleep(singleTestDuration)
					}},
					testing.InternalTest{F: func(*testing.T) {
						time.Sleep(singleTestDuration)
					}},
				)
			})
		})
		elapsed := time.Since(start)
		if elapsed < singleTestDuration {
			t.Errorf("tests execution time %s not exceeded %s", elapsed, singleTestDuration)
		}
		if elapsed > expectedMaxDuration {
			t.Errorf("tests execution time %s exceeded %s", elapsed, expectedMaxDuration)
		}
	})
	t.Run("should run 2 suites parallel in one test", func(t *testing.T) {
		singleTestDuration := 1 * time.Second
		expectedMaxDuration := singleTestDuration + 300*time.Millisecond
		start := time.Now()
		t.Run("internal", func(it *testing.T) {
			// internal2 is necessary because PackageParallel calls t.Parallel()
			it.Run("internal2", func(it2 *testing.T) {
				pt.PackageParallel(it2, testing.InternalTest{F: func(*testing.T) {
					time.Sleep(singleTestDuration)
				}})
				pt.PackageParallel(it2, testing.InternalTest{F: func(*testing.T) {
					time.Sleep(singleTestDuration)
				}})
			})
		})
		elapsed := time.Since(start)
		if elapsed < singleTestDuration {
			t.Errorf("tests execution time %s not exceeded %s", elapsed, singleTestDuration)
		}
		if elapsed > expectedMaxDuration {
			t.Errorf("tests execution time %s exceeded %s", elapsed, expectedMaxDuration)
		}
	})
	t.Run("should run 2 suites parallel in different tests", func(t *testing.T) {
		singleTestDuration := 1 * time.Second
		expectedMaxDuration := singleTestDuration + 300*time.Millisecond
		start := time.Now()
		t.Run("internal", func(it *testing.T) {
			// internal2 is necessary because PackageParallel calls t.Parallel()
			it.Run("internal2", func(it2 *testing.T) {
				pt.PackageParallel(it2, testing.InternalTest{F: func(*testing.T) {
					time.Sleep(singleTestDuration)
				}})
				pt.PackageParallel(it2, testing.InternalTest{F: func(*testing.T) {
					time.Sleep(singleTestDuration)
				}})
			})
			it.Run("internal3", func(it3 *testing.T) {
				pt.PackageParallel(it3, testing.InternalTest{F: func(*testing.T) {
					time.Sleep(singleTestDuration)
				}})
				pt.PackageParallel(it3, testing.InternalTest{F: func(*testing.T) {
					time.Sleep(singleTestDuration)
				}})
			})
		})
		elapsed := time.Since(start)
		if elapsed < singleTestDuration {
			t.Errorf("tests execution time %s not exceeded %s", elapsed, singleTestDuration)
		}
		if elapsed > expectedMaxDuration {
			t.Errorf("tests execution time %s exceeded %s", elapsed, expectedMaxDuration)
		}
	})
}

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
		pt.Parallel(t, testing.InternalTest{F: func(*testing.T) {}})
	})
	t.Run("should not panic with 2 tests", func(t *testing.T) {
		t.Parallel()
		pt.Parallel(t,
			testing.InternalTest{F: func(*testing.T) {}},
			testing.InternalTest{F: func(*testing.T) {}},
		)
	})
	t.Run("should run 1 test", func(t *testing.T) {
		t.Parallel()
		called := false
		t.Run("internal", func(it *testing.T) {
			pt.Parallel(it, testing.InternalTest{F: func(*testing.T) {
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
	t.Run("should run 2 tests", func(t *testing.T) {
		t.Parallel()
		called1 := false
		called2 := false
		t.Run("internal", func(it *testing.T) {
			pt.Parallel(it,
				testing.InternalTest{F: func(*testing.T) {
					if called1 {
						t.Error("test1 is called twice")
					}
					called1 = true
				}},
				testing.InternalTest{F: func(*testing.T) {
					if called2 {
						t.Error("test2 is called twice")
					}
					called2 = true
				}},
			)
		})
		if !called1 {
			t.Error("test1 is not called")
		}
		if !called2 {
			t.Error("test2 is not called")
		}
	})
}

func TestParallel2(t *testing.T) {
	// Do not call t.Parallel() because test measures execution time
	t.Run("should run 2 tests parallel", func(t *testing.T) {
		singleTestDuration := 1 * time.Second
		expectedMaxDuration := singleTestDuration + 300*time.Millisecond
		start := time.Now()
		t.Run("internal", func(it *testing.T) {
			pt.Parallel(it,
				testing.InternalTest{F: func(*testing.T) {
					time.Sleep(singleTestDuration)
				}},
				testing.InternalTest{F: func(*testing.T) {
					time.Sleep(singleTestDuration)
				}},
			)
		})
		elapsed := time.Since(start)
		if elapsed < singleTestDuration {
			t.Errorf("tests execution time %s not exceeded %s", elapsed, singleTestDuration)
		}
		if elapsed > expectedMaxDuration {
			t.Errorf("tests execution time %s exceeded %s", elapsed, expectedMaxDuration)
		}
	})
	t.Run("should run 10 tests parallel", func(t *testing.T) {
		singleTestDuration := 1 * time.Second
		expectedMaxDuration := singleTestDuration * 10 / 2
		tests := make([]testing.InternalTest, 10)
		for i := range tests {
			tests[i] = testing.InternalTest{F: func(*testing.T) {
				time.Sleep(singleTestDuration)
			}}
		}
		start := time.Now()
		t.Run("internal", func(it *testing.T) {
			pt.Parallel(it, tests...)
		})
		elapsed := time.Since(start)
		if elapsed < singleTestDuration {
			t.Errorf("tests execution time %s not exceeded %s", elapsed, singleTestDuration)
		}
		if elapsed > expectedMaxDuration {
			t.Errorf("tests execution time %s exceeded %s", elapsed, expectedMaxDuration)
		}
	})
	t.Run("should run 2 suites parallel", func(t *testing.T) {
		singleTestDuration := 1 * time.Second
		expectedMaxDuration := singleTestDuration + 300*time.Millisecond
		start := time.Now()
		t.Run("internal", func(it *testing.T) {
			pt.Parallel(it, testing.InternalTest{F: func(*testing.T) {
				time.Sleep(singleTestDuration)
			}})
			pt.Parallel(it, testing.InternalTest{F: func(*testing.T) {
				time.Sleep(singleTestDuration)
			}})
		})
		elapsed := time.Since(start)
		if elapsed < singleTestDuration {
			t.Errorf("tests execution time %s not exceeded %s", elapsed, singleTestDuration)
		}
		if elapsed > expectedMaxDuration {
			t.Errorf("tests execution time %s exceeded %s", elapsed, expectedMaxDuration)
		}
	})
	t.Run("should run 10 suites parallel", func(t *testing.T) {
		singleTestDuration := 1 * time.Second
		expectedMaxDuration := singleTestDuration * 10 / 2
		start := time.Now()
		t.Run("internal", func(it *testing.T) {
			for i := 0; i < 10; i++ {
				pt.Parallel(it, testing.InternalTest{F: func(*testing.T) {
					time.Sleep(singleTestDuration)
				}})
			}
		})
		elapsed := time.Since(start)
		if elapsed < singleTestDuration {
			t.Errorf("tests execution time %s not exceeded %s", elapsed, singleTestDuration)
		}
		if elapsed > expectedMaxDuration {
			t.Errorf("tests execution time %s exceeded %s", elapsed, expectedMaxDuration)
		}
	})
	t.Run("should run 2 suites sequential", func(t *testing.T) {
		singleTestDuration := 1 * time.Second
		expectedMinDuration := 2 * singleTestDuration
		expectedMaxDuration := 3 * singleTestDuration
		start := time.Now()
		t.Run("internal", func(it *testing.T) {
			it.Run("internal2", func(it2 *testing.T) {
				pt.Parallel(it2, testing.InternalTest{F: func(*testing.T) {
					time.Sleep(singleTestDuration)
				}})
			})
			it.Run("internal2", func(it2 *testing.T) {
				pt.Parallel(it2, testing.InternalTest{F: func(*testing.T) {
					time.Sleep(singleTestDuration)
				}})
			})
		})
		elapsed := time.Since(start)
		if elapsed < expectedMinDuration {
			t.Errorf("tests execution time %s not exceeded %s", elapsed, expectedMinDuration)
		}
		if elapsed > expectedMaxDuration {
			t.Errorf("tests execution time %s exceeded %s", elapsed, expectedMaxDuration)
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
		internalTest := pt.Group("", testing.InternalTest{F: func(*testing.T) {
			if called {
				t.Error("test is called twice")
			}
			called = true
		}})
		t.Run("internal", func(it *testing.T) {
			internalTest.F(it)
		})
		if !called {
			t.Error("test is not called")
		}
	})
	t.Run("should run 2 tests", func(t *testing.T) {
		t.Parallel()
		called1 := false
		called2 := false
		internalTest := pt.Group("",
			testing.InternalTest{F: func(*testing.T) {
				if called1 {
					t.Error("test1 is called twice")
				}
				called1 = true
			}},
			testing.InternalTest{F: func(*testing.T) {
				if called2 {
					t.Error("test2 is called twice")
				}
				called2 = true
			}},
		)
		t.Run("internal", func(it *testing.T) {
			internalTest.F(it)
		})
		if !called1 {
			t.Error("test1 is not called")
		}
		if !called2 {
			t.Error("test2 is not called")
		}
	})
}

func TestGroup2(t *testing.T) {
	// Do not call t.Parallel() because test measures execution time
	t.Run("should run 2 tests parallel", func(t *testing.T) {
		singleTestDuration := 1 * time.Second
		expectedMaxDuration := singleTestDuration + 300*time.Millisecond
		internalTest := pt.Group("",
			testing.InternalTest{F: func(*testing.T) {
				time.Sleep(singleTestDuration)
			}},
			testing.InternalTest{F: func(*testing.T) {
				time.Sleep(singleTestDuration)
			}},
		)
		start := time.Now()
		t.Run("internal", func(it *testing.T) {
			internalTest.F(it)
		})
		elapsed := time.Since(start)
		if elapsed < singleTestDuration {
			t.Errorf("tests execution time %s not exceeded %s", elapsed, singleTestDuration)
		}
		if elapsed > expectedMaxDuration {
			t.Errorf("tests execution time %s exceeded %s", elapsed, expectedMaxDuration)
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
		internalTest := pt.Test("abc", func(*testing.T) {})
		if internalTest.Name != "abc" {
			t.Error("name is wrong")
		}
	})
	t.Run("should return right test func", func(t *testing.T) {
		t.Parallel()
		called := false
		testFunc := func(*testing.T) {
			if called {
				t.Error("test is called twice")
			}
			called = true
		}
		internalTest := pt.Test("", testFunc)
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
