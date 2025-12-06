package aufgabe6

import (
	"fmt"
)

func ExampleIndexOfSmallest() {
	fmt.Println(IndexOfSmallest([]int{4, 2, 8, 1}))
	fmt.Println(IndexOfSmallest([]int{7}))
	fmt.Println(IndexOfSmallest([]int{-5, -2, 0}))

	// Output:
	// 3
	// 0
	// 0
}
