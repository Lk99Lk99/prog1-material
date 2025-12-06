package aufgabe3

/* AUFGABENSTELLUNG: Vervollständigen Sie die unten stehende Funktion.
 * RANDBEDINGUNG: Die Funktion muss rekursiv sein.
 * ERREICHBARE PUNKTE: 10
 */

// SumList liefert die Summe aller Zahlen in list.
func SumList(list []int) int {
	// TODO
	if len(list) == 0 {
		return 0
	}
	return SumList(list[1:]) + list[0]
}
