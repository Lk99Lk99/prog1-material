package aufgabe1

/* AUFGABENSTELLUNG: Vervollständigen Sie die unten stehende Funktion.
 * ERREICHBARE PUNKTE: 10
 */

// PrefixBelow10 erwartet eine Liste "list" von Zahlen und liefert
// die längste Teil-Liste, mit der "list" beginnt und die nur Zahlen < 10 enthält.
func PrefixBelow10(list []int) []int {
	// TODO

	ziel := []int{}
	if len(list) == 0 {
		return ziel
	}
	if list[0] < 10 {
		for i := 0; i < len(list); i++ {
			if list[i] >= 10 {
				return ziel
			}
			if list[i] < 10 {
				ziel = append(ziel, list[i])
			}
		}
	}
	return ziel
}
