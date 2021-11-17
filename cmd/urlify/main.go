package main

import (
	"fmt"

	"github.com/egorovdmi/go-challenges/internals/platform/strings"
)

func main() {
	urlify("Mr John Smith    ", 13)
}

func urlify(s string, realLen int) {
	println("#================================================================")
	result := strings.URLify(s, realLen)
	println(s, result)

	println("#================================================================")
	result2 := strings.URLify2("Mr John Smith")
	fmt.Printf("%s %s\n", s, string(result2))

	println("#================================================================")
	data := []rune(s)
	strings.URLify3(data, realLen)
	fmt.Printf("%s %s\n", s, string(data))
}
