package aufgabe4

import "fmt"

func ExampleRepeatStrings() {
	fmt.Println(RepeatStrings([]string{"a", "b", "c"}))
	fmt.Println(RepeatStrings([]string{"hi"}))
	fmt.Println(RepeatStrings([]string{}))

	// Output:
	// [a a b b c c]
	// [hi hi]
	// []
}
