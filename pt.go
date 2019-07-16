package pt

import (
	"reflect"
	"testing"
)

func PackageParallel(t *testing.T, tests ...testing.InternalTest) {
	if t == nil {
		panic("argument t *testing.T can not be nil")
	}
	if !alreadyParallel(t) { // avoid panic
		t.Parallel()
	}
	Parallel(t, tests...)
}

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

func Group(name string, tests ...testing.InternalTest) testing.InternalTest {
	return testing.InternalTest{
		Name: name,
		F: func(t *testing.T) {
			Parallel(t, tests...)
		},
	}
}

func Test(name string, test func(t *testing.T)) testing.InternalTest {
	if test == nil {
		panic("argument test func(t *testing.T) can not be nil")
	}
	return testing.InternalTest{
		Name: name,
		F:    test,
	}
}

func alreadyParallel(t *testing.T) bool {
	// copy of mutex is not used, so can ignore govet error
	testObject := reflect.ValueOf(*t) // nolint:govet
	isParallelField := testObject.FieldByName("isParallel")
	isParallel := isParallelField.Bool()
	return isParallel
}
