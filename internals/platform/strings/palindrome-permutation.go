package strings

import "strings"

func IsPalindromePermutation(str string) (bool, []string) {
	dict := make(map[rune]int)
	length := len(str)

	for _, c := range strings.ToLower(str) {
		dict[c]++
	}

	oddCount := 0
	var oddKey rune

	for key, value := range dict {
		if value%2 > 0 && key != ' ' {
			if oddCount > 0 {
				return false, []string{}
			}

			oddCount++
			oddKey = key
		}
	}

	runes := make([]rune, length)
	cursor := 0
	leftCursor := length - 1
	for key, value := range dict {
		n := value / 2
		maxCursor := cursor + n
		for cursor < maxCursor {
			runes[cursor] = key
			runes[leftCursor] = key
			cursor += 1
			leftCursor -= 1
		}

		// for spaces
		if value%2 > 0 {
			if key == ' ' {
				runes[cursor] = key
				cursor++
			}
		}
	}

	runes[cursor] = oddKey

	return true, []string{string(runes)}
}
