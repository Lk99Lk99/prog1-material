package aufgabe5

// IsChain erwartet eine Liste von Dominoe-Objekten.
// Die Funktion prüft, ob diese Liste eine Kette bildet,
// die nach den Domino-Regeln erlaubt wäre.
// Bei einer solchen Kette muss immer die rechte Seite eines Steins
// gleich der linken Seite des nächsten Steins sein.
func IsChain(dominoes []Dominoe) bool {

	for i := 1; i < len(dominoes); i++ {
		if dominoes[i].Left != dominoes[i-1].Right {
			return false
		}
	}
	return true
}

// Dominoe repräsentiert einen Domino-Stein mit zwei Zahlen.
type Dominoe struct {
	Left  int
	Right int
}
