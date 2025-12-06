package aufgabe4

// ElementProducts erwartet zwei int-Listen l1 und l2.
// Sie liefert eine Liste, die an jeder Position
// jeweils das Produkt der beiden Elemente enthält.
// Falls eine Position nur in einer Liste vorkommt,
// soll dieses Element ins Ergebnis übernommen werden.
func ElementProducts(l1, l2 []int) []int {
	result := []int{}
	max := len(l1)
	if max < len(l2) {
		max = len(l2)
	}

	for i := 0; i < max; i++ {
		dif1 := 1
		dif2 := 1
		if i < len(l1) {
			dif1 = l1[i]
		}
		if i < len(l2) {
			dif2 = l2[i]
		}
		result = append(result, (dif1 * dif2))
	}

	return result
}
