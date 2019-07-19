package example

import (
	"testing"

	"github.com/maratori/pt"
)

func TestSum(t *testing.T) {
	pt.PackageParallel(t,
		pt.Test("should be 0 without values", func(t *testing.T) {
			if sum() != 0 {
				t.Fail()
			}
		}),
		pt.Group("should be equal to single value",
			pt.Test("0", testSumSingleValue(0)),
			pt.Test("1", testSumSingleValue(1)),
			pt.Test("-1", testSumSingleValue(-1)),
			pt.Test("123", testSumSingleValue(123)),
			pt.Test("-123", testSumSingleValue(-123)),
		),
		pt.Group("should be sum of two values",
			pt.Test("0+0 = 0", testSumTwoValues(0, 0, 0)),
			pt.Test("0+1 = 1", testSumTwoValues(0, 1, 1)),
			pt.Test("1+0 = 1", testSumTwoValues(1, 0, 1)),
			pt.Test("5+6 = 11", testSumTwoValues(5, 6, 11)),
		),
		pt.Test("1+2+3+4+5 = 15", func(t *testing.T) {
			if sum(1, 2, 3, 4, 5) != 15 {
				t.Fail()
			}
		}),
	)
}

func TestFibonacci(t *testing.T) {
	pt.PackageParallel(t,
		pt.Test("0 -> 0", testFibonacci(0, 0)),
		pt.Test("1 -> 1", testFibonacci(1, 1)),
		pt.Test("2 -> 1", testFibonacci(2, 1)),
		pt.Test("3 -> 2", testFibonacci(3, 2)),
		pt.Test("4 -> 3", testFibonacci(4, 3)),
	)
}

func testSumSingleValue(value int) func(t *testing.T) {
	return func(t *testing.T) {
		if sum(value) != value {
			t.Fail()
		}
	}
}

func testSumTwoValues(a int, b int, expected int) func(t *testing.T) {
	return func(t *testing.T) {
		if sum(a, b) != expected {
			t.Fail()
		}
	}
}

func testFibonacci(n uint64, expected uint64) func(t *testing.T) {
	return func(t *testing.T) {
		if fibonacci(n) != expected {
			t.Fail()
		}
	}
}
