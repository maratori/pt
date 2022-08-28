//go:build example
// +build example

package example

func sum(numbers ...int) int {
	s := 0
	for _, n := range numbers {
		s += n
	}
	return s
}

func fibonacci(n uint64) uint64 {
	if n <= 1 {
		return n
	}

	var n2, n1 uint64 = 0, 1

	for i := uint64(2); i < n; i++ {
		n2, n1 = n1, n1+n2
	}

	return n2 + n1
}
