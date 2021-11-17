package strings

func URLify(s string, realLength int) string {
	runes := []rune(s)
	cursor := len(runes) - 1
	for i := realLength - 1; i >= 0; i-- {
		if runes[i] != ' ' {
			runes[cursor] = runes[i]
			cursor--
		} else {
			runes[cursor] = '0'
			runes[cursor-1] = '2'
			runes[cursor-2] = '%'
			cursor -= 3
		}
	}

	return string(runes)
}

func URLify2(s string) []rune {
	spacesCount := countSymbol(s, ' ')
	runes := make([]rune, len(s)+spacesCount*2)
	cursor := 0
	input := []rune(s)
	for i := 0; i < len(input); i++ {
		if input[i] != ' ' {
			runes[cursor] = input[i]
			cursor++
		} else {
			runes[cursor] = '%'
			runes[cursor+1] = '2'
			runes[cursor+2] = '0'
			cursor += 3
		}
	}

	return runes
}

func URLify3(data []rune, realLength int) {
	cursor := len(data) - 1
	for i := realLength - 1; i >= 0; i-- {
		if data[i] != ' ' {
			data[cursor] = data[i]
			cursor--
		} else {
			data[cursor] = '0'
			data[cursor-1] = '2'
			data[cursor-2] = '%'
			cursor -= 3
		}
	}
}

func URLify4(s string) []rune {
	var runes []rune
	for _, r := range s {
		if r != ' ' {
			runes = append(runes, r)
		} else {
			runes = append(runes, '%', '2', '0')
		}
	}

	return runes
}

func countSymbol(s string, symbol rune) int {
	count := 0
	for _, r := range s {
		if r == symbol {
			count++
		}
	}

	return count
}
