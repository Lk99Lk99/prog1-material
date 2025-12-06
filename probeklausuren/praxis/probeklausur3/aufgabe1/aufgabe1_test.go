package aufgabe1

import "fmt"

func ExampleCountDivisible() {
	fmt.Println(CountDivisible([]int{3, 6, 9, 10}, 3))
	fmt.Println(CountDivisible([]int{2, 4, 6, 8}, 2))
	fmt.Println(CountDivisible([]int{1, 3, 5}, 2))
	fmt.Println(CountDivisible([]int{}, 5))

	// Output:
	// 3
	// 4
	// 0
	// 0
}
