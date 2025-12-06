package aufgabe3

import "math"

/* AUFGABENSTELLUNG: Vervollständigen Sie die unten stehende Funktion.
 * RANDBEDINGUNG: Die Funktion muss rekursiv sein.
 * ERREICHBARE PUNKTE: 10
 */

// CountSquares erwartet eine Liste von Zahlen.
// CountSquares liefert die Anzahl der QuadratzahlenZahlen in der Liste.

func CountSquares(list []int) int {
	// TODO

	if len(list) == 0 {
		return +0
	}

	if IsSquare(float64(list[0])) {
		return CountSquares(list[1:]) + 1
	}

	return CountSquares(list[1:])
}

func IsSquare(x float64) bool {

	return math.Sqrt(x)*math.Sqrt(x) == x

}
