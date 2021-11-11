package strings

func CheckPermutations(s1 string, s2 string) bool {
	m1 := toMap(s1)
	m2 := toMap(s2)

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

func toMap(s string) map[rune]int {
	m := make(map[rune]int)
	for _, r := range s {
		m[r]++
	}

	return m
}
