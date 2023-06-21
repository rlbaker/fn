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
