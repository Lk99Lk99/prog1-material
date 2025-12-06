package aufgabe4

/* AUFGABENSTELLUNG: Vervollständigen Sie die unten stehende Funktion.
 * ERREICHBARE PUNKTE: 10
 */

// ElementSums erwartet zwei int-Listen l1 und l2.
// Sie liefert eine Liste, die an jeder Position
// jeweils die Summe der beiden Elemente enthält.
//
// Annahmen für die Berechnung:
// Falls eine Liste kürzer ist als die andere, soll für die Berechnung der
// hinteren Werte ihr letztes Element verwendet werden.
// Für leere Listen soll für die Berechnung ggf. 0 verwendet werden.
func ElementSums(l1, l2 []int) []int {
	result := []int{}
	// TODO

	/*if len(l1) == 0 {
		return l2
	}

	if len(l2) == 0 {
		return l1
	}

	if len(l1) < len(l2) {
		dif := len(l2) - len(l1)
		for i := 0; i < len(l1); i++ {
			sum := l1[i] + l2[i]
			result = append(result, sum)

		}
		for j := len(l1); j < len(l2); j++ {
			sum := l1[dif] + l2[j]
			result = append(result, sum)
		}
		return result
	}

	if len(l2) < len(l1) {
		dif := len(l1) - len(l2)
		for i := 0; i < len(l2); i++ {
			sum := l1[i] + l2[i]
			result = append(result, sum)

		}
		for j := len(l2); j < len(l1); j++ {
			sum := l2[dif] + l1[j]
			result = append(result, sum)
		}
		return result
	}

	for i := 0; i < len(l1); i++ {
		sum := l1[i] + l2[i]
		result = append(result, sum)
	}*/

	max := len(l1)
	if max < len(l2) {
		max = len(l2)
	}

	current1 := 0
	current2 := 0
	for i := 0; i < max; i++ {
		if i < len(l1) {
			current1 = l1[i]
		}

		if i < len(l2) {
			current2 = l2[i]
		}

		result = append(result, (current1 + current2))
	}

	return result
}
