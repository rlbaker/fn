# fn

The `fn` package implements a variety functional programming helpers using Go generics. All functions are implemented using the provided `Reduce` function.

## Functions

### `Reduce`: Applies a function against an accumulator and each element in the array (from left to right) to reduce it to a single output value.

```go
input := []int{1, 2, 3, 4, 5}
result := fn.Reduce(input, 0, func(acc int, curr int) int {
	return acc + curr
})
fmt.Println(result) // 15
```

### `Fold`: Reduces a list of items into a single output, by applying a binary function to pairs of items, starting from the first item and the second, until the end of the list.


```go
input := []int{1, 2, 3, 4, 5}
result := fn.Fold(input, func(acc int, curr int) int {
	return acc * curr
})
fmt.Println(result) // 120
```

### `Map`: Creates a new array populated with the results of calling a provided function on every element in the input array.

```go
input := []int{1, 2, 3, 4, 5}
result := fn.Map(input, func(i int) int {
	return i * i
})
fmt.Println(result) // [1 4 9 16 25]
```

### `FlatMap`: Similar to `Map`, but each input item can be mapped to zero or more output items (so `func` should return a slice of zero or more items).

```go
input := []int{1, 2, 3}
result := fn.FlatMap(input, func(i int) []int {
	return []int{i, i * i}
})
fmt.Println(result) // [1 1 2 4 3 9]
```

### `Filter`: Creates a new array with all elements that pass the test implemented by the provided function.

```go
input := []int{1, 2, 3, 4, 5}
result := fn.Filter(input, func(i int) bool {
	return i%2 == 0
})
fmt.Println(result) // [2 4]
```

### `Each`: Executes a provided function once for each array element.

```go
input := []int{1, 2, 3, 4, 5}
fn.Each(input, func(i int) {
	fmt.Println(i)
})
```

### `Some`: Tests whether at least one element in the array passes the test implemented by the provided function.


```go
input := []int{1, 2, 3, 4, 5}
result := fn.Some(input, func(i int) bool {
	return i > 4
})
fmt.Println(result) // true
```

### `Every`: Tests whether all elements in the array pass the test implemented by the provided function.

```go
input := []int{1, 2, 3, 4, 5}
result := fn.Every(input, func(i int) bool {
	return i < 6
})
fmt.Println(result) // true
```

### `Count`: Counts the number of elements in the array that pass the test implemented by the provided function.

```go
input := []int{1, 2, 3, 4, 5}
result := fn.Count(input, func(i int) bool {
	return i%2 == 0
})
fmt.Println(result) // 2
```

### `CompactNil`: Removes all nil values from the slice.

```go
input := []*int{nil, new(int), nil, new(int)}
result := fn.CompactNil(input)
fmt.Println(len(result)) // 2
```

### `Flatten`: Flattens a nested slice (the slice of slices) into a single slice.

```go
input := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
result := fn.Flatten(input)
fmt.Println(result) // [1 2 3 4 5 6 7 8 9]
```
