package main

import "github.com/egorovdmi/cs50-assignments/internals/platform/strings"

func main() {
	sb := strings.NewStringBuilder()
	sb.Append("hello \n")
	sb.Append("world")

	println(sb.ToString())
}
