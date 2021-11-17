package strings

func OneAway(strA string, strB string) bool {
	lenA := len(strA)
	lenB := len(strB)

	// potentailly more then one char inserted
	if abs(lenA-lenB) > 1 {
		return false
	}

	// if symbol was replaced
	if lenA == lenB {
		runesB := []rune(strB)
		notEqualSymbols := 0
		for i, s := range strA {
			if s != runesB[i] {
				if notEqualSymbols > 0 {
					return false
				}
				notEqualSymbols++
			}
		}

		// if strings are equal => false
		return notEqualSymbols != 0
	}

	// if only one symbol difference
	if lenA > lenB {
		return oneSymbolDiff(strA, strB)
	} else {
		return oneSymbolDiff(strB, strA)
	}
}

func oneSymbolDiff(longStr string, shortStr string) bool {
	symDiffCount := 0
	cursor := 0
	runes := []rune(shortStr)
	shortLen := len(shortStr)
	for _, s := range longStr {
		if cursor < shortLen && runes[cursor] != s {
			if symDiffCount > 0 {
				return false
			}

			symDiffCount++
		} else {
			cursor++
		}
	}

	return true
}

func abs(n int) int {
	if n < 0 {
		return n * -1
	}

	return n
}
