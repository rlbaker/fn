package fn

import "fmt"

func ExampleReduce() {
	input := []int{1, 2, 3, 4, 5}

	result := Reduce(input, 0, func(acc int, curr int) int {
		return acc + curr
	})

	fmt.Println(result)
	// Output: 15
}

func ExampleFold() {
	input := []int{1, 2, 3, 4, 5}

	result := Fold(input, func(acc int, curr int) int {
		return acc * curr
	})

	fmt.Println(result)
	// Output: 120
}

func ExampleMap() {
	input := []int{1, 2, 3, 4, 5}

	result := Map(input, func(i int) int {
		return i * i
	})

	fmt.Println(result)
	// Output: [1 4 9 16 25]
}

func ExampleFlatMap() {
	input := []int{1, 2, 3}

	result := FlatMap(input, func(i int) []int {
		return []int{i, i * i}
	})

	fmt.Println(result)
	// Output: [1 1 2 4 3 9]
}

func ExampleFilter() {
	input := []int{1, 2, 3, 4, 5}

	result := Filter(input, func(i int) bool {
		return i%2 == 0
	})

	fmt.Println(result)
	// Output: [2 4]
}

func ExampleEach() {
	input := []int{1, 2, 3, 4, 5}
	Each(input, func(i int) {
		fmt.Println(i)
	})

	// Output: 1
	// 2
	// 3
	// 4
	// 5
}

func ExampleSome() {
	input := []int{1, 2, 3, 4, 5}

	result := Some(input, func(i int) bool {
		return i > 4
	})

	fmt.Println(result)
	// Output: true
}

func ExampleEvery() {
	input := []int{1, 2, 3, 4, 5}

	result := Every(input, func(i int) bool {
		return i < 6
	})

	fmt.Println(result)
	// Output: true
}

func ExampleCount() {
	input := []int{1, 2, 3, 4, 5}

	result := Count(input, func(i int) bool {
		return i%2 == 0
	})

	fmt.Println(result)
	// Output: 2
}

func ExampleCompactNil() {
	input := []*int{nil, new(int), nil, new(int)}

	result := CompactNil(input)

	fmt.Println(len(result))
	// Output: 2
}

func ExampleFlatten() {
	input := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}

	result := Flatten(input)

	fmt.Println(result)
    // Output: [1 2 3 4 5 6 7 8 9]
}
