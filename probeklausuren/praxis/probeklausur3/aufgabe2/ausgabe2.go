package aufgabe2

/* AUFGABENSTELLUNG: Vervollständigen Sie unten stehende Funktion.
 * ERREICHBARE PUNKTE: 10
 */

// RemoveSpaces erwartet einen String s und liefert s zurück,
// allerdings ohne Leerzeichen. Andere Zeichen bleiben unverändert.
func RemoveSpaces(s string) string {
	result := ""
	// TODO
	for i := 0; i < len(s); i++ {
		if s[i] != ' ' {
			result = result + string(s[i])
		}
	}
	return result
}
