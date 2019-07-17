/*
Package pt provides functions to run tests in parallel.
You don't have to call t.Parallel() anymore.


Example

	func TestMyFunc(t *testing.T) {
		pt.PackageParallel(t,
			pt.Test("should do something", func(t *testing.T) {
				// test code
			}),
			pt.Test("should do something else", func(t *testing.T) {
				// test code
			}),
			pt.Group("some condition",
				pt.Test("should not do something", func(t *testing.T) {
					// test code
				}),
				pt.Test("should not do something else", func(t *testing.T) {
					// test code
				}),
			}),
		)
	}

You can achieve the same behavior with bare testing package:

	func TestMyFunc(t *testing.T) {
		t.Parallel()
		t.Run("should do something", func(t *testing.T) {
			t.Parallel()
			// test code
		})
		t.Run("should do something else", func(t *testing.T) {
			t.Parallel()
			// test code
		})
		t.Run("some condition", func(t *testing.T) {
			t.Parallel()
			t.Run("should not do something", func(t *testing.T) {
				t.Parallel()
				// test code
			})
			t.Run("should not do something else", func(t *testing.T) {
				t.Parallel()
				// test code
			})
		})
	}
*/
package pt

import (
	"reflect"
	"testing"
)

/*
PackageParallel runs provided tests in parallel with other tests in package.
It is designed to be used with Group and Test.

	func TestA(t *testing.T) {
		pt.PackageParallel(t, test1, test2)
	}

is equivalent to

	func TestA(t *testing.T) {
		t.Parallel()
		t.Run(test1.Name, func(t *testing.T) {
			t.Parallel()
			test1.F(t)
		})
		t.Run(test2.Name, func(t *testing.T) {
			t.Parallel()
			test2.F(t)
		})
	}

If you don't need to run TestA in parallel with other tests, use Parallel.
*/
func PackageParallel(t *testing.T, tests ...testing.InternalTest) {
	if t == nil {
		panic("argument t *testing.T can not be nil")
	}
	if !alreadyParallel(t) { // avoid panic
		t.Parallel()
	}
	Parallel(t, tests...)
}

/*
Parallel runs provided tests in parallel.
It is designed to be used with Group and Test.

	func TestA(t *testing.T) {
		pt.Parallel(t, test1, test2)
	}

is equivalent to

	func TestA(t *testing.T) {
		t.Run(test1.Name, func(t *testing.T) {
			t.Parallel()
			test1.F(t)
		})
		t.Run(test2.Name, func(t *testing.T) {
			t.Parallel()
			test2.F(t)
		})
	}


Note that Parallel will not execute different TestXXX and TestYYY in parallel.
For example test3 and test4 will run in parallel and after that test5 and test6 will run in parallel.

	func TestB(t *testing.T) {
		pt.Parallel(t, test3, test4)
	}
	func TestC(t *testing.T) {
		pt.Parallel(t, test5, test6)
	}

If you need different behavior, use PackageParallel.
*/
func Parallel(t *testing.T, tests ...testing.InternalTest) {
	if t == nil {
		panic("argument t *testing.T can not be nil")
	}
	for _, test := range tests {
		test := test
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()
			test.F(t)
		})
	}
}

// Group is a constructor of testing.InternalTest.
// It wraps provided tests with a single testing.InternalTest.
// Provided tests will run in parallel when wrapper is executed.
// It is designed to be used as an argument of Group, Parallel and PackageParallel.
func Group(name string, tests ...testing.InternalTest) testing.InternalTest {
	return testing.InternalTest{
		Name: name,
		F: func(t *testing.T) {
			Parallel(t, tests...)
		},
	}
}

// Test is a simple constructor of testing.InternalTest.
// It is designed to be used as an argument of Group, Parallel and PackageParallel.
func Test(name string, test func(t *testing.T)) testing.InternalTest {
	if test == nil {
		panic("argument test func(t *testing.T) can not be nil")
	}
	return testing.InternalTest{
		Name: name,
		F:    test,
	}
}

// alreadyParallel returns value of private field isParallel for provided t *testing.T
func alreadyParallel(t *testing.T) bool {
	// copy of mutex is not used, so can ignore govet error
	testObject := reflect.ValueOf(*t) // nolint:govet
	isParallelField := testObject.FieldByName("isParallel")
	isParallel := isParallelField.Bool()
	return isParallel
}
