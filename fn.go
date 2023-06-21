package fn

var none struct{}

func Reduce[T, U any](input []T, init U, reducefn func(acc U, curr T) U) U {
	acc := init
	for i := range input {
		acc = reducefn(acc, input[i])
	}
	return acc
}

func Fold[T, U any](input []T, foldf func(acc T, curr T) T) T {
	return Reduce(input[1:], input[0], func(acc T, curr T) T {
		return foldf(acc, curr)
	})
}

func Map[T, U any](input []T, mapf func(T) U) []U {
	return Reduce(input, []U{}, func(acc []U, curr T) []U {
		return append(acc, mapf(curr))
	})
}

func FlatMap[T, U any](input []T, mapf func(T) []U) []U {
	return Reduce(input, []U{}, func(acc []U, curr T) []U {
		return append(acc, mapf(curr)...)
	})
}

func Filter[T any](input []T, filterf func(T) bool) []T {
	return Reduce(input, []T{}, func(acc []T, curr T) []T {
		if filterf(curr) {
			return append(acc, curr)
		}
		return acc
	})
}

func Each[T any](input []T, eachf func(T)) {
	Reduce(input, none, func(acc struct{}, curr T) struct{} {
		eachf(curr)
		return acc
	})
}

func Some[T any](input []T, somef func(T) bool) bool {
	return Reduce(input, false, func(acc bool, curr T) bool {
		return acc || somef(curr)
	})
}

func Every[T any](input []T, everyf func(T) bool) bool {
	return Reduce(input, true, func(acc bool, curr T) bool {
		return acc && everyf(curr)
	})
}

func Count[T any](input []T, countf func(T) bool) int {
	return Reduce(input, 0, func(acc int, curr T) int {
		if countf(curr) {
			return acc + 1
		}
		return acc
	})
}

func CompactNil[T any](input []*T) []*T {
	return Reduce(input, []*T{}, func(acc []*T, curr *T) []*T {
		if curr != nil {
			return append(acc, curr)
		}
		return acc
	})
}

func Flatten[T any](input [][]T) []T {
	return Reduce(input, []T{}, func(acc []T, curr []T) []T {
		return append(acc, curr...)
	})
}
