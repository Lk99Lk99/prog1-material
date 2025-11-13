package chesspiece

import "fmt"

func ExampleChessPiece_MoveAllowed() {
	l1 := ChessPiece{
		pieceType: BISHOP,
		colour:    WHITE,
		row:       2,
		column:    2,
	}

	l2 := ChessPiece{
		pieceType: KNIGHT,
		colour:    WHITE,
		row:       2,
		column:    6,
	}

	l3 := ChessPiece{
		pieceType: ROOK,
		colour:    WHITE,
		row:       4,
		column:    4,
	}

	l4 := ChessPiece{
		pieceType: QUEEN,
		colour:    WHITE,
		row:       4,
		column:    4,
	}

	fmt.Println(l1.MoveAllowed(3, 4))
	fmt.Println(l1.MoveAllowed(0, 0))
	fmt.Println("KNIGHT")
	fmt.Println(l2.MoveAllowed(3, 4))
	fmt.Println(l2.MoveAllowed(0, 0))
	fmt.Println(l2.MoveAllowed(4, 5))
	fmt.Println(l2.MoveAllowed(3, 8))
	fmt.Println(l2.MoveAllowed(0, 5))
	fmt.Println(l2.MoveAllowed(0, 7))
	fmt.Println("ROOK")
	fmt.Println(l3.MoveAllowed(3, 4))
	fmt.Println(l3.MoveAllowed(1, 4))
	fmt.Println(l3.MoveAllowed(4, 5))
	fmt.Println(l3.MoveAllowed(4, 8))
	fmt.Println(l3.MoveAllowed(0, 5))
	fmt.Println(l3.MoveAllowed(3, 3))
	fmt.Println("QUEEN")
	fmt.Println(l4.MoveAllowed(3, 4))
	fmt.Println(l4.MoveAllowed(1, 4))
	fmt.Println(l4.MoveAllowed(4, 5))
	fmt.Println(l4.MoveAllowed(4, 8))
	fmt.Println(l4.MoveAllowed(5, 5))
	fmt.Println(l4.MoveAllowed(8, 8))
	fmt.Println(l4.MoveAllowed(1, 1))
	fmt.Println(l4.MoveAllowed(0, 0))
	fmt.Println(l4.MoveAllowed(6, 5))
	fmt.Println(l4.MoveAllowed(3, 2))
	fmt.Println(l4.MoveAllowed(1, 5))
	fmt.Println(l4.MoveAllowed(8, 9))

	// Output:
	// false
	// true
}
