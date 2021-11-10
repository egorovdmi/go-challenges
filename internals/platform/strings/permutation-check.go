package strings

func CheckPermutations(s1 string, s2 string) bool {
	m1 := toMap(toRunesSlice(s1))
	m2 := toMap(toRunesSlice(s2))

	for k := range m1 {
		if m1[k] != m2[k] {
			return false
		}
	}

	for k := range m2 {
		if m1[k] != m2[k] {
			return false
		}
	}

	return true
}

func toMap(runes []rune) map[rune]int {
	m := make(map[rune]int)
	for _, r := range runes {
		m[r]++
	}

	return m
}

func toRunesSlice(s string) []rune {
	var runes []rune
	for _, r := range s {
		runes = append(runes, r)
	}

	return runes
}
