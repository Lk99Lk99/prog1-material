package aufgabe5

import "fmt"

func ExampleExists() {
	dict := []Entry{{"Hund", "dog"}, {"Katze", "cat"}}

	fmt.Println(Exists(dict, "Hund"))
	fmt.Println(Exists(dict, "Baum"))
	fmt.Println(Exists(dict, ""))

	// Output:
	// true
	// false
	// false
}
