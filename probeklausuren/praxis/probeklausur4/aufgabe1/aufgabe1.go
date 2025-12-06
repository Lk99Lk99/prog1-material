package aufgabe1

import "strings"

// LongestAbc erwartet eine Liste von Strings und liefert
// das längste Element, das mit der Buchstabenfolge "abc" beginnt.
// Liefert den leeren String, falls es kein solches Element gibt.
func LongestAbc(list []string) string {
	longest := ""
	for i := 0; i < len(list); i++ {
		if strings.HasPrefix(list[i], "abc") && list[i] > longest {
			longest = list[i]
		}
	}

	return longest
}
