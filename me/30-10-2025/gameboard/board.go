package gameboard

import (
	"fmt"
	"strings"
)

// Verwende eine string-Matrix für das Spielfeld.

// Newboard liefert eine 3x3-Matrix aus Strings.
func NewBoard() [][]string {
	board := [][]string{}

	// Füge drei Zeilen ui board hinzu.
	for row := 0; row < 3; row++ {
		board = append(board, []string{" ", " ", " "})
	}

	return board
}

// PrintBoard gibt ein Spielfeld menschenlesbar auf der Konsole aus.
func PrintBoard(b [][]string) {
	div := strings.Repeat("+---", len(b[0])) + "+"
	for row := 0; row < len(b); row++ {
		fmt.Println(div)
		for col := 0; col < len(b[row]); col++ {
			fmt.Printf("| %v ", b[row][col])
		}
		fmt.Println("|")
	}
	fmt.Println(div)
}
