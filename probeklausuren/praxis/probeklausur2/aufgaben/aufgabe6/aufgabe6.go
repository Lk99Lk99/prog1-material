package aufgabe6

/* AUFGABENSTELLUNG: Vervollständigen Sie die unten stehende Funktion.
 * ERREICHBARE PUNKTE: 10
 */

// DuplicateSinglets erwartet eine int-Liste list.
// Die Funktion liefert eine int-Liste, bei der alle Elemente,
// die in list nur einmal vorkommen, verdoppelt sind,
// also zwei Mal hintereinander stehen.
// Elemente, die schon in list mehrfach vorkommen, sollen wie sie sind
// ins Ergebnis übertragen werden.
func DuplicateSinglets(list []int) []int {
	result := []int{}
	// TODO
	for i := range list {

		if !IsSingle(list, list[i]) {
			result = append(result, list[i], list[i])
		} else {
			result = append(result, list[i])
		}

	}

	return result
}

func IsSingle(list []int, n int) bool {
	count := 0
	for i := range list {
		if list[i] == n {
			count++
		}
	}
	return count >= 2
}
