package aufgabe6

/* AUFGABENSTELLUNG: Vervollständigen Sie die unten stehende Funktion.
 * ERREICHBARE PUNKTE: 10
 */

// IndexOfSmallest erwartet eine Liste list.
// Die Funktion liefert den Index der kleinsten Zahl.
// Es kann angenommen werden, list enthält mindestens ein Element.
func IndexOfSmallest(list []int) int {
	// TODO

	min := list[0]
	index := 0
	for i, el := range list {
		if el < min {
			min = el
			index = i
		}
	}

	return index
}
