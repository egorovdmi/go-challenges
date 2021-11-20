package strings

import "strconv"

func CompressString(str string) string {
	result := make([]rune, 0, len(str))
	count := 0
	var prevoiusSymbol rune

	for i, s := range str {
		if len(result) > len(str) {
			return str
		}

		if i == 0 { // first symbol
			result = append(result, s)
			prevoiusSymbol = s
			count++
		} else if prevoiusSymbol == s { // the same symbol
			count++
		} else if prevoiusSymbol != s { // symbol has changed
			result = append(result, []rune(strconv.FormatInt(int64(count), 10))...)
			result = append(result, s)
			prevoiusSymbol = s
			count = 1
		}
	}

	result = append(result, []rune(strconv.FormatInt(int64(count), 10))...)

	if len(result) > len(str) {
		return str
	} else {
		return string(result)
	}
}
