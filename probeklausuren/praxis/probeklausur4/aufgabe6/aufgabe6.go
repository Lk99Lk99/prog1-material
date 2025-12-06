package aufgabe6

// import "slices"

// Duplicates erwartet eine int-Liste list.
// Die Funktion liefert eine int-Liste mit den Elementen,
// die mehr als einmal in l1 vorkommen.
// Im Ergebnis kommt jedes Element nur einmal vor.
// Die Elemente stehen im Ergebnis in der Reihenfolge ihres ersten Auftretens.
func Duplicates(list []int) []int {
	result := []int{}

	for i := 0; i < len(list); i++ {
		if IsDouble(list, list[i]) && !IsDouble(result, list[i]) {
			result = append(result, list[i])
		}
	}

	return result
}

func IsDouble(list []int, x int) bool {
	count := 0
	for i := 0; i < len(list); i++ {
		if list[i] == x {
			count++
		}
	}
	return count < 2
}
