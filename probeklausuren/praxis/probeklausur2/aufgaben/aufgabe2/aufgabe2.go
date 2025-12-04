package aufgabe2

/* AUFGABENSTELLUNG: Vervollständigen Sie unten stehende Funktion.
 * ERREICHBARE PUNKTE: 10
 */

// FilterDigits erwartet einen String s und liefert einen String zurück,
// der aus s entsteht, indem alle Ziffern entfernt werden.
// Alle anderen Zeichen sollen unverändert bleiben.
func FilterDigits(s string) string {
	result := ""
	// TODO

	for i := 0; i < len(s); i++ {
		//	if s[i] == 0 ||s[i] == 1 ||s[i] == 2 ||s[i] == 3 ||s[i] == 4 ||s[i] == 5 ||s[i] == 6 ||s[i] == 7||s[i] == 8 ||s[i] == 9 {

		//	}
		result = append(result, s[i])

		//}
	}

	return result
}
