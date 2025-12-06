package aufgabe5

/* AUFGABENSTELLUNG: Vervollständigen Sie die unten stehende Funktion,
 * die auf dem darunter definierten Datentyp arbeitet.
 * MAX. PUNKTE: 10
 */

// Exists prüft, ob in dict ein Eintrag für de existiert.
// Falls ja, soll true zurückgegeben werden, sonst false.
func Exists(dict []Entry, de string) bool {
	// TODO
	for i := 0; i < len(dict); i++ {
		if dict[i].De == de || dict[i].En == de {
			return true
		}
	}
	return false
}

// Entry repräsentiert einen Eintrag in einem Wörterbuch
type Entry struct {
	De string
	En string
}
