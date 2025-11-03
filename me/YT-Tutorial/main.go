package main

import (
	"fmt"
	"prog1-material/me/YT-Tutorial/calculator"
)

func main() {
	fmt.Println("23 + 42 =", calculator.Add(23, 42))
	fmt.Println(calculator.Divide(17, 3))
	fmt.Println(calculator.Divide(20, 5))
}
