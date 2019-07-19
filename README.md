# <img src="logo.png" height="100px" alt="Logo"> <br> [![Build Status](https://travis-ci.com/maratori/pt.svg?branch=master)](https://travis-ci.com/maratori/pt) [![codecov](https://codecov.io/gh/maratori/pt/branch/master/graph/badge.svg)](https://codecov.io/gh/maratori/pt) [![codebeat badge](https://codebeat.co/badges/60157255-e2dd-4819-a0c5-4ac164f57b88)](https://codebeat.co/projects/github-com-maratori-pt-master) [![Maintainability](https://api.codeclimate.com/v1/badges/0078c4d48b975f84c1c9/maintainability)](https://codeclimate.com/github/maratori/pt/maintainability) [![Go Report Card](https://goreportcard.com/badge/github.com/maratori/pt)](https://goreportcard.com/report/github.com/maratori/pt) [![GitHub](https://img.shields.io/github/license/maratori/pt.svg)](LICENSE) [![GoDoc](https://godoc.org/github.com/maratori/pt?status.svg)](http://godoc.org/github.com/maratori/pt)


This is a go (golang) package with functions to **P**arallel **T**ests run.
You don't have to call `t.Parallel()` anymore.


## Installation

```bash
go get github.com/maratori/pt
```
or
```bash
dep ensure -add github.com/maratori/pt
```


## Usage

You can use `pt.PackageParallel`, `pt.Parallel`, `pt.Group`, `pt.Test` in standard go test function to run tests parallel.

See [example_test.go](example/example_test.go)

```go
...

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

...
```

When you run go test, all these tests run in parallel.

```bash
go test -v github.com/maratori/pt/example
```

<details><summary>Output</summary>

```
=== RUN   TestSum
=== PAUSE TestSum
=== RUN   TestFibonacci
=== PAUSE TestFibonacci
=== CONT  TestSum
=== CONT  TestFibonacci
=== RUN   TestFibonacci/0_->_0
=== RUN   TestSum/should_be_0_without_values
=== PAUSE TestFibonacci/0_->_0
=== RUN   TestFibonacci/1_->_1
=== PAUSE TestFibonacci/1_->_1
=== RUN   TestFibonacci/2_->_1
=== PAUSE TestFibonacci/2_->_1
=== RUN   TestFibonacci/3_->_2
=== PAUSE TestFibonacci/3_->_2
=== PAUSE TestSum/should_be_0_without_values
=== RUN   TestFibonacci/4_->_3
=== PAUSE TestFibonacci/4_->_3
=== RUN   TestSum/should_be_equal_to_single_value
=== CONT  TestFibonacci/2_->_1
=== PAUSE TestSum/should_be_equal_to_single_value
=== CONT  TestFibonacci/0_->_0
=== CONT  TestFibonacci/1_->_1
=== CONT  TestFibonacci/3_->_2
=== CONT  TestFibonacci/4_->_3
=== RUN   TestSum/should_be_sum_of_two_values
=== PAUSE TestSum/should_be_sum_of_two_values
=== RUN   TestSum/1+2+3+4+5_=_15
=== PAUSE TestSum/1+2+3+4+5_=_15
=== CONT  TestSum/should_be_0_without_values
=== CONT  TestSum/should_be_sum_of_two_values
=== RUN   TestSum/should_be_sum_of_two_values/0+0_=_0
=== PAUSE TestSum/should_be_sum_of_two_values/0+0_=_0
=== RUN   TestSum/should_be_sum_of_two_values/0+1_=_1
--- PASS: TestFibonacci (0.00s)
    --- PASS: TestFibonacci/0_->_0 (0.00s)
    --- PASS: TestFibonacci/2_->_1 (0.00s)
    --- PASS: TestFibonacci/1_->_1 (0.00s)
    --- PASS: TestFibonacci/3_->_2 (0.00s)
    --- PASS: TestFibonacci/4_->_3 (0.00s)
=== CONT  TestSum/1+2+3+4+5_=_15
=== CONT  TestSum/should_be_equal_to_single_value
=== RUN   TestSum/should_be_equal_to_single_value/0
=== PAUSE TestSum/should_be_sum_of_two_values/0+1_=_1
=== PAUSE TestSum/should_be_equal_to_single_value/0
=== RUN   TestSum/should_be_sum_of_two_values/1+0_=_1
=== RUN   TestSum/should_be_equal_to_single_value/1
=== PAUSE TestSum/should_be_sum_of_two_values/1+0_=_1
=== PAUSE TestSum/should_be_equal_to_single_value/1
=== RUN   TestSum/should_be_sum_of_two_values/5+6_=_11
=== PAUSE TestSum/should_be_sum_of_two_values/5+6_=_11
=== CONT  TestSum/should_be_sum_of_two_values/0+0_=_0
=== CONT  TestSum/should_be_sum_of_two_values/5+6_=_11
=== CONT  TestSum/should_be_sum_of_two_values/0+1_=_1
=== RUN   TestSum/should_be_equal_to_single_value/-1
=== CONT  TestSum/should_be_sum_of_two_values/1+0_=_1
=== PAUSE TestSum/should_be_equal_to_single_value/-1
=== RUN   TestSum/should_be_equal_to_single_value/123
=== PAUSE TestSum/should_be_equal_to_single_value/123
=== RUN   TestSum/should_be_equal_to_single_value/-123
=== PAUSE TestSum/should_be_equal_to_single_value/-123
=== CONT  TestSum/should_be_equal_to_single_value/0
=== CONT  TestSum/should_be_equal_to_single_value/-1
=== CONT  TestSum/should_be_equal_to_single_value/123
=== CONT  TestSum/should_be_equal_to_single_value/-123
=== CONT  TestSum/should_be_equal_to_single_value/1
--- PASS: TestSum (0.00s)
    --- PASS: TestSum/should_be_0_without_values (0.00s)
    --- PASS: TestSum/1+2+3+4+5_=_15 (0.00s)
    --- PASS: TestSum/should_be_sum_of_two_values (0.00s)
        --- PASS: TestSum/should_be_sum_of_two_values/0+0_=_0 (0.00s)
        --- PASS: TestSum/should_be_sum_of_two_values/0+1_=_1 (0.00s)
        --- PASS: TestSum/should_be_sum_of_two_values/5+6_=_11 (0.00s)
        --- PASS: TestSum/should_be_sum_of_two_values/1+0_=_1 (0.00s)
    --- PASS: TestSum/should_be_equal_to_single_value (0.00s)
        --- PASS: TestSum/should_be_equal_to_single_value/0 (0.00s)
        --- PASS: TestSum/should_be_equal_to_single_value/123 (0.00s)
        --- PASS: TestSum/should_be_equal_to_single_value/-1 (0.00s)
        --- PASS: TestSum/should_be_equal_to_single_value/-123 (0.00s)
        --- PASS: TestSum/should_be_equal_to_single_value/1 (0.00s)
PASS
ok      github.com/maratori/pt/example  0.006s
```
</details>


## Flags for `go test`

There are 2 flags for `go test` command related to parallel run.

* -parallel n
> Allow parallel execution of test functions that call t.Parallel.
> The value of this flag is the maximum number of tests to run
> simultaneously; by default, it is set to the value of GOMAXPROCS.
> Note that -parallel only applies within a single test binary.
> The 'go test' command may run tests for different packages
> in parallel as well, according to the setting of the -p flag
> (see 'go help build').

* -p n
> The number of programs, such as build commands or
> test binaries, that can be run in parallel.
> The default is the number of CPUs available.

So `go test -p 2 -parallel 1` will run tests from two packages in parallel, but does not parallel tests within that packages.  
On the other hand `go test -p 1 -parallel 2` will run tests from different packages sequentially. And run two tests in parallel within single package.


## Difference between `pt.PackageParallel` and `pt.Parallel`

The difference can be demonstrated with code below.  
Tests will run in parallel: 1-8.  
After that tests will run in parallel: 9-12.  
After that tests will run in parallel: 13-16.  
See [godoc](https://godoc.org/github.com/maratori/pt) for more info.  

```go
func TestA(t *testing.T) {
	pt.PackageParallel(t, test1, test2)
	pt.PackageParallel(t, test3, test4)
}
func TestB(t *testing.T) {
	pt.PackageParallel(t, test5, test6)
	pt.Parallel(t, test7, test8)
}
func TestC(t *testing.T) {
	pt.Parallel(t, test9, test10)
	pt.Parallel(t, test11, test12)
}
func TestD(t *testing.T) {
	pt.Parallel(t, test13, test14)
	pt.Parallel(t, test15, test16)
}
```

## Changelog

### [v1.0.0] - 2019-07-21

#### Added
* Functions: PackageParallel, Parallel, Group, Test
* [GoDoc](http://godoc.org/github.com/maratori/pt)
* Example [package](example)
* Unit [tests](pt_test.go) for all functions
* All possible linters in [golangci-lint](https://github.com/golangci/golangci-lint) ([config](.golangci.yml))
* Project [logo](logo.png)
* MIT [license](LICENSE)

## License

[MIT License](LICENSE)
