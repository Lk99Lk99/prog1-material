package fizzbuzz

import "fmt"

// Fizz gibt die Zahlen von 1 bis 30 aus und ersetz dabei jede
// durch 3 teilbare Zahl durch "fizz" und jede durch 5 teilbare
// Zahl durch "buzz". Bei Zahlen, die sowohl durch 3 als auch durch
// 5 teilbar sind, wird "fizzbuzz" ausgegeben.
func FizzImproved(x, y, n int) {
	for i := 1; i <= n; i++ {
		// wenn i durch 3 und 5 teilbar ist, gib "fizzbuzz" aus
		if i%x == 0 && i%y == 0 {

			fmt.Println("fizzbuzz")
		} else if
		// wenn i durch 3 teilbar ist, gib "fizz" aus
		i%x == 0 {
			fmt.Println("fizz")
		} else if
		// wenn i durch 5 teilbar ist, gib "buzz" aus
		i%y == 0 {
			fmt.Println("buzz")
		} else {
			// sonst gib i aus
			fmt.Println(i)
		}
	}
}
