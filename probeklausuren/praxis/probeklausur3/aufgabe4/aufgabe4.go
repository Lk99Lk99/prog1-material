package aufgabe4

/* AUFGABENSTELLUNG: Vervollständigen Sie die unten stehende Funktion.
 * ERREICHBARE PUNKTE: 10
 */

// RepeatStrings erwartet eine Liste von Strings words.
// Die Funktion liefert eine neue Liste,
// in der jedes Wort doppelt hintereinander steht.
func RepeatStrings(words []string) []string {
	result := []string{}
	// TODO
	for i := 0; i < len(words); i++ {
		result = append(result, words[i], words[i])
	}

	return result
}
