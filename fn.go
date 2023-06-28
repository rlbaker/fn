package fn

var none struct{}

// Fold applies a function to an accumulator and each element in the slice (from left to right) to produce a single output value.
func Fold[T, U any](input []T, init U, foldf func(acc U, curr T) U) U {
	acc := init
	for i := range input {
		acc = foldf(acc, input[i])
	}
	return acc
}

// Reduce folds a slice into a single output, by applying a binary function to pairs of elements, starting from the first element and the second, until the end of the slice.
func Reduce[T any](input []T, reducef func(acc T, curr T) T) T {
	return Fold(input[1:], input[0], func(acc T, curr T) T {
		return reducef(acc, curr)
	})
}

// Map creates a new slice populated with the results of calling a provided function on every element in the input slice.
func Map[T, U any](input []T, mapf func(T) U) []U {
	return Fold(input, []U{}, func(acc []U, curr T) []U {
		return append(acc, mapf(curr))
	})
}

// FlatMap is similar to Map, but each input element can be mapped to zero or more output elements, each of which will be appended to the result slice.
func FlatMap[T, U any](input []T, mapf func(T) []U) []U {
	return Fold(input, []U{}, func(acc []U, curr T) []U {
		return append(acc, mapf(curr)...)
	})
}

// Filter creates a new slice with all elements that pass the test implemented by the provided function.
func Filter[T any](input []T, filterf func(T) bool) []T {
	return Fold(input, []T{}, func(acc []T, curr T) []T {
		if filterf(curr) {
			return append(acc, curr)
		}
		return acc
	})
}

// Each executes a provided function once for each slice element.
func Each[T any](input []T, eachf func(T)) {
	Fold(input, none, func(acc struct{}, curr T) struct{} {
		eachf(curr)
		return acc
	})
}

// Some tests whether at least one element in the slice passes the test implemented by the provided function.
func Some[T any](input []T, somef func(T) bool) bool {
	return Fold(input, false, func(acc bool, curr T) bool {
		return acc || somef(curr)
	})
}

// Every tests whether all elements in the slice pass the test implemented by the provided function.
func Every[T any](input []T, everyf func(T) bool) bool {
	return Fold(input, true, func(acc bool, curr T) bool {
		return acc && everyf(curr)
	})
}

// Count counts the number of elements in the slice that pass the test implemented by the provided function.
func Count[T any](input []T, countf func(T) bool) int {
	return Fold(input, 0, func(acc int, curr T) int {
		if countf(curr) {
			return acc + 1
		}
		return acc
	})
}

// CompactNil removes all nil values from the slice.
func CompactNil[T any](input []*T) []*T {
	return Fold(input, []*T{}, func(acc []*T, curr *T) []*T {
		if curr != nil {
			return append(acc, curr)
		}
		return acc
	})
}

// Flatten flattens a nested slice (slice of slices) into a single slice.
func Flatten[T any](input [][]T) []T {
	return Fold(input, []T{}, func(acc []T, curr []T) []T {
		return append(acc, curr...)
	})
}
