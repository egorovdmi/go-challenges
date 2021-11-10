package main

import "github.com/egorovdmi/cs50-assignments/internals/platform/trie"

func main() {
	t := trie.NewTrie()
	t.AddWord("apple")

	r := t.HasWord("applee")
	println(r)
}
