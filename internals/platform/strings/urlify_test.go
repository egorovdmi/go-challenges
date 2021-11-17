package strings_test

import (
	"testing"

	"github.com/egorovdmi/go-challenges/internals/platform/strings"
)

func BenchmarkURLify(b *testing.B) {
	s := "Mr John Smith    "

	for n := 0; n < b.N; n++ {
		strings.URLify(s, 13)
	}
}

func BenchmarkURLify2(b *testing.B) {
	s := `Mr John SmithMr John SmithMr John SmithMr John SmithMr John SmithMr John SmithMr John Smith
	Mr John SmithMr John SmithMr John SmithMr John SmithMr John SmithMr John SmithMr John Smith
	Mr John SmithMr John SmithMr John SmithMr John SmithMr John SmithMr John SmithMr John Smith
	Mr John SmithMr John SmithMr John SmithMr John SmithMr John SmithMr John SmithMr John Smith`

	for n := 0; n < b.N; n++ {
		strings.URLify2(s)
	}
}

var result3 []rune

func BenchmarkURLify3(b *testing.B) {
	s := "Mr John Smith    "
	var data []rune

	for n := 0; n < b.N; n++ {
		data = []rune(s)
		strings.URLify3(data, 13)
	}

	result3 = data
}

func BenchmarkURLify4(b *testing.B) {
	s := `Mr John SmithMr John SmithMr John SmithMr John SmithMr John SmithMr John SmithMr John Smith
	Mr John SmithMr John SmithMr John SmithMr John SmithMr John SmithMr John SmithMr John Smith
	Mr John SmithMr John SmithMr John SmithMr John SmithMr John SmithMr John SmithMr John Smith
	Mr John SmithMr John SmithMr John SmithMr John SmithMr John SmithMr John SmithMr John Smith`

	for n := 0; n < b.N; n++ {
		strings.URLify4(s)
	}
}
