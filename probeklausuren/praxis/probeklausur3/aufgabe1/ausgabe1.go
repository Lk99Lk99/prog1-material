package aufgabe1

/* AUFGABENSTELLUNG: Vervollständigen Sie die unten stehende Funktion.
 * ERREICHBARE PUNKTE: 10
 */

// CountDivisible erwartet eine Liste list und eine Zahl k.
// Die Funktion liefert die Anzahl der Werte in list,
// die ohne Rest durch k teilbar sind.
func CountDivisible(list []int, k int) int {
	// TODO
	if len(list) == 0 {
		return 0
	}
	if list[0]%k == 0 {
		return CountDivisible(list[1:], k) + 1
	}
	return CountDivisible(list[1:], k)
}
