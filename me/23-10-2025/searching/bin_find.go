package searching

// / Liefert die Position von x in der Liste l,
// / oder liefert -1, falls x nicht in l vorkommt.
func BinFind1(l []int, x int) int {
	links := 0
	for len(l) > 0 {
		mitte := len(l) / 2
		// Vergleiche x mit dem Element in der Mitte.
		// Wenn x == l[mitte], dann fertig.
		if x == l[mitte] {
			return mitte + links
		}
		// Wenn x < l[mitte], dann suche im linken Teil weiter.
		if x < l[mitte] {
			// Lasse nur den linken Teil der Liste übrig.
			// Alles von 0 bis exklusive Mitte.
			// l = l[0:mitte]
			l = l[:mitte]
		} else {
			// Wenn x > l[mitte], dann suche im rechten Teil weiter.
			// Lasse nur den rechten Teil der Liste übrig.
			// l = l[mitte+1 : len(l)]
			l = l[mitte+1:]
			links += mitte + 1

		}

	}
	return -1
}

func BinFind2(l []int, x int) int {
	links := 0       // Hier beginnt der interessante Teil der Liste.
	rechts := len(l) // Hier endet der interessante Teil der Liste.
	for rechts > links {
		mitte := (rechts-links)/2 + links

		if x == l[mitte] {
			return mitte
		}
		if x < l[mitte] {
			rechts = mitte
		} else {
			links = mitte + 1
		}
	}
	return -1
}

func BinFindRek(l []int, x int) int {
	if len(l) == 0 {
		return -1
	}

	mitte := len(l) / 2

	if x == l[mitte] {
		return mitte
	}
	if x < l[mitte] {
		// Suche im linken Teil weiter
		return BinFindRek(l[:mitte], x)
	}
	if x > l[mitte] {
		// Suche im rechten Teil weiter
		result := BinFindRek(l[mitte+1:], x)
		if result == -1 {
			return -1
		}
		return result + mitte + 1
	}
	return -1
}
