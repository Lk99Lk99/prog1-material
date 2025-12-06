package aufgabe2

import "fmt"

func ExampleRemoveSpaces() {
	fmt.Println(RemoveSpaces("a b c"))
	fmt.Println(RemoveSpaces("hello world"))
	fmt.Println(RemoveSpaces("nospaces"))
	fmt.Println(RemoveSpaces(""))
	fmt.Println(RemoveSpaces("  a  "))

	// Output:
	// abc
	// helloworld
	// nospaces
	//
	// a
}
