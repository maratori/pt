# <img src="logo.png" height="100px" alt="Logo"> <br> [![CI][ci-img]][ci-url] [![Codecov][codecov-img]][codecov-url] [![Codebeat][codebeat-img]][codebeat-url] [![Maintainability][codeclimate-img]][codeclimate-url] [![Go Report Card][goreportcard-img]][goreportcard-url] [![License][license-img]][license-url] [![Go Reference][godoc-img]][godoc-url]


This is a go (golang) package with functions to run **P**arallel **T**ests.
You don't have to call `t.Parallel()` anymore.


## Installation

```bash
go get github.com/maratori/pt
```


## Usage

You can use `pt.PackageParallel`, `pt.Parallel`, `pt.Group`, `pt.Test` in standard go test function to run tests in parallel.

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

When you call `go test`, all these tests will run in parallel.

```bash
go test -v -p 8 -parallel 8 --tags example github.com/maratori/pt/example
```

<details><summary>Output</summary>

```
=== RUN   TestSum
=== PAUSE TestSum
=== RUN   TestFibonacci
=== PAUSE TestFibonacci
=== CONT  TestSum
=== RUN   TestSum/should_be_0_without_values
=== PAUSE TestSum/should_be_0_without_values
=== RUN   TestSum/should_be_equal_to_single_value
=== PAUSE TestSum/should_be_equal_to_single_value
=== CONT  TestFibonacci
=== RUN   TestFibonacci/0_->_0
=== RUN   TestSum/should_be_sum_of_two_values
=== PAUSE TestSum/should_be_sum_of_two_values
=== RUN   TestSum/1+2+3+4+5_=_15
=== PAUSE TestFibonacci/0_->_0
=== PAUSE TestSum/1+2+3+4+5_=_15
=== CONT  TestSum/should_be_0_without_values
=== RUN   TestFibonacci/1_->_1
=== PAUSE TestFibonacci/1_->_1
=== RUN   TestFibonacci/2_->_1
=== CONT  TestSum/1+2+3+4+5_=_15
=== PAUSE TestFibonacci/2_->_1
=== RUN   TestFibonacci/3_->_2
=== PAUSE TestFibonacci/3_->_2
=== CONT  TestSum/should_be_equal_to_single_value
=== RUN   TestFibonacci/4_->_3
=== PAUSE TestFibonacci/4_->_3
=== CONT  TestFibonacci/0_->_0
=== CONT  TestSum/should_be_sum_of_two_values
=== RUN   TestSum/should_be_sum_of_two_values/0+0_=_0
=== CONT  TestFibonacci/2_->_1
=== PAUSE TestSum/should_be_sum_of_two_values/0+0_=_0
=== CONT  TestFibonacci/4_->_3
=== CONT  TestFibonacci/1_->_1
=== CONT  TestFibonacci/3_->_2
=== RUN   TestSum/should_be_equal_to_single_value/0
=== PAUSE TestSum/should_be_equal_to_single_value/0
=== RUN   TestSum/should_be_sum_of_two_values/0+1_=_1
=== RUN   TestSum/should_be_equal_to_single_value/1
=== PAUSE TestSum/should_be_sum_of_two_values/0+1_=_1
=== PAUSE TestSum/should_be_equal_to_single_value/1
--- PASS: TestFibonacci (0.00s)
    --- PASS: TestFibonacci/0_->_0 (0.00s)
    --- PASS: TestFibonacci/2_->_1 (0.00s)
    --- PASS: TestFibonacci/4_->_3 (0.00s)
    --- PASS: TestFibonacci/1_->_1 (0.00s)
    --- PASS: TestFibonacci/3_->_2 (0.00s)
=== RUN   TestSum/should_be_equal_to_single_value/-1
=== PAUSE TestSum/should_be_equal_to_single_value/-1
=== RUN   TestSum/should_be_sum_of_two_values/1+0_=_1
=== RUN   TestSum/should_be_equal_to_single_value/123
=== PAUSE TestSum/should_be_sum_of_two_values/1+0_=_1
=== PAUSE TestSum/should_be_equal_to_single_value/123
=== RUN   TestSum/should_be_sum_of_two_values/5+6_=_11
=== RUN   TestSum/should_be_equal_to_single_value/-123
=== PAUSE TestSum/should_be_equal_to_single_value/-123
=== PAUSE TestSum/should_be_sum_of_two_values/5+6_=_11
=== CONT  TestSum/should_be_sum_of_two_values/0+0_=_0
=== CONT  TestSum/should_be_equal_to_single_value/0
=== CONT  TestSum/should_be_equal_to_single_value/123
=== CONT  TestSum/should_be_sum_of_two_values/1+0_=_1
=== CONT  TestSum/should_be_equal_to_single_value/-123
=== CONT  TestSum/should_be_equal_to_single_value/1
=== CONT  TestSum/should_be_sum_of_two_values/0+1_=_1
=== CONT  TestSum/should_be_sum_of_two_values/5+6_=_11
=== CONT  TestSum/should_be_equal_to_single_value/-1
--- PASS: TestSum (0.00s)
    --- PASS: TestSum/should_be_0_without_values (0.00s)
    --- PASS: TestSum/1+2+3+4+5_=_15 (0.00s)
    --- PASS: TestSum/should_be_sum_of_two_values (0.00s)
        --- PASS: TestSum/should_be_sum_of_two_values/0+0_=_0 (0.00s)
        --- PASS: TestSum/should_be_sum_of_two_values/1+0_=_1 (0.00s)
        --- PASS: TestSum/should_be_sum_of_two_values/0+1_=_1 (0.00s)
        --- PASS: TestSum/should_be_sum_of_two_values/5+6_=_11 (0.00s)
    --- PASS: TestSum/should_be_equal_to_single_value (0.00s)
        --- PASS: TestSum/should_be_equal_to_single_value/0 (0.00s)
        --- PASS: TestSum/should_be_equal_to_single_value/123 (0.00s)
        --- PASS: TestSum/should_be_equal_to_single_value/-123 (0.00s)
        --- PASS: TestSum/should_be_equal_to_single_value/1 (0.00s)
        --- PASS: TestSum/should_be_equal_to_single_value/-1 (0.00s)
PASS
ok  	github.com/maratori/pt/example	0.091s
```
</details>

## Supported golang versions

* 1.8
* 1.9
* 1.10
* 1.11
* 1.12
* 1.13
* 1.14
* 1.15
* 1.16
* 1.17
* 1.18
* 1.19


## Flags for `go test`

There are 2 flags for `go test` command related to parallel execution.

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

So `go test -p 8 -parallel 1` runs test packages in parallel, while tests inside each package are executed sequentially.  
On the other hand, `go test -p 1 -parallel 8` runs different packages sequentially, but tests inside each package are executed in parallel.


## Difference between `pt.PackageParallel` and `pt.Parallel`

The difference can be demonstrated with the code below. Tests will be executed in the following sequence:
1. Tests 1-8 run in parallel
1. After that tests 9-12 run in parallel
1. After that tests 13-16 run in parallel

See [godoc][godoc-url] for more info.  

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

### [v1.0.2] - 2022-08-28

#### Changed
* Use go 1.19
* Test all supported go versions on CI (1.8 .. 1.19)

#### Fixed
* Follow new go 1.19 doc comments [conventions](https://go.dev/doc/comment)

### [v1.0.1] - 2019-09-29

#### Changed
* Use go 1.13

#### Fixed
* Readme and godoc slightly improved

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

[MIT License][license-url]



[ci-img]: https://github.com/maratori/pt/actions/workflows/ci.yml/badge.svg
[ci-url]: https://github.com/maratori/pt/actions/workflows/ci.yml
[codecov-img]: https://codecov.io/gh/maratori/pt/branch/main/graph/badge.svg?token=WisNd8SOoW
[codecov-url]: https://codecov.io/gh/maratori/pt
[codebeat-img]: https://codebeat.co/badges/95684dfc-294c-4712-a20b-fb19c6e6b0c5
[codebeat-url]: https://codebeat.co/projects/github-com-maratori-pt-main
[codeclimate-img]: https://api.codeclimate.com/v1/badges/0078c4d48b975f84c1c9/maintainability
[codeclimate-url]: https://codeclimate.com/github/maratori/pt/maintainability
[goreportcard-img]: https://goreportcard.com/badge/github.com/maratori/pt
[goreportcard-url]: https://goreportcard.com/report/github.com/maratori/pt
[license-img]: https://img.shields.io/github/license/maratori/pt.svg
[license-url]: https://github.com/maratori/pt/blob/main/LICENSE
[godoc-img]: https://pkg.go.dev/badge/github.com/maratori/pt.svg
[godoc-url]: https://pkg.go.dev/github.com/maratori/pt
