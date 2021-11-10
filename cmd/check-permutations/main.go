package main

import "github.com/egorovdmi/go-challenges/internals/platform/strings"

func main() {
	checkPermutations("hello", "oellh")
	checkPermutations("hello1", "oellh")
	checkPermutations("go-challenges", "challenges-go")
	checkPermutations("это челлендж", "челлендж это")
	checkPermutations("go", "Go")
}

func checkPermutations(s1 string, s2 string) {
	println("#================================================================")
	if strings.CheckPermutations(s1, s2) {
		println("characters in srings: '", s1, "' and '", s2, "': SAME.")
	} else {
		println("characters in srings: '", s1, "' and '", s2, "': DIFFERENT.")
	}
}
