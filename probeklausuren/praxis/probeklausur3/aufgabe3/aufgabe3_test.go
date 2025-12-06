package aufgabe3

import "fmt"

func ExampleSumList() {
	fmt.Println(SumList([]int{1, 2, 3, 4}))
	fmt.Println(SumList([]int{5}))
	fmt.Println(SumList([]int{}))

	// Output:
	// 10
	// 5
	// 0
}
