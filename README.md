# fn

[![Go Reference](https://pkg.go.dev/badge/github.com/rlbaker/fn.svg)](https://pkg.go.dev/github.com/rlbaker/fn)

The `fn` package implements a variety functional programming helpers using Go generics.
All functions are implemented using the provided `Fold` function.

You probably don't need this package.
In most situations a `for` loop will be more idiomatic and consistent with other Go code.
This package was created for the purpose of examining some of the functional programming patterns that _could_ be replicated using Go generics.

## Usage

Add package dependency: `go get -u github.com/rlbaker/fn@v0.1.1`

```go
package main

import "github.com/rlbaker/fn"

func main() {
	input := []int{1, 2, 3, 4, 5}

	folded := Fold(input, 0, func(acc int, curr int) int {
		return acc + curr
	})
	fmt.Println(folded) // 15

	mapped := Map(input, func(x int) int {
		return x * x
	})
	fmt.Println(mapped) // [1 4 9 16 25]
}
```

## Available Functions

- `Fold` applies a function to an accumulator and each element in the slice (from left to right) to produce a single output value.
- `Reduce` folds a slice into a single output, by applying a binary function to pairs of elements, starting from the first element and the second, until the end of the slice.
- `Map` creates a new slice populated with the results of calling a provided function on every element in the input slice.
- `FlatMap` is similar to `Map`, but each input element can be mapped to zero or more output elements, each of which will be appended to the result slice.
- `Filter` creates a new slice with all elements that pass the test implemented by the provided function.
- `Each` executes a provided function once for each slice element.
- `Some` tests whether at least one element in the slice passes the test implemented by the provided function.
- `Every` tests whether all elements in the slice pass the test implemented by the provided function.
- `Count` counts the number of elements in the slice that pass the test implemented by the provided function.
- `CompactNil` removes all nil values from the slice.
- `Flatten` flattens a nested slice (slice of slices) into a single slice.
