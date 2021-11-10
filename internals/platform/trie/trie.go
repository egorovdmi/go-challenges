package trie

import "fmt"

const size = 28

type trieNode struct {
	terminatesHere bool
	nodes          [size]*trieNode
}

func newTrieNode() *trieNode {
	return &trieNode{}
}

type Trie struct {
	nodes [size]*trieNode
}

func NewTrie() *Trie {
	var t Trie
	return &t
}

func (t *Trie) AddWord(word string) {
	var currentNode *trieNode

	for _, l := range word {
		if currentNode == nil && t.nodes[index(l)] == nil {
			t.nodes[index(l)] = newTrieNode()
			currentNode = t.nodes[index(l)]
		} else if currentNode.nodes[index(l)] == nil {
			currentNode.nodes[index(l)] = newTrieNode()
			currentNode = currentNode.nodes[index(l)]
		}
	}

	currentNode.terminatesHere = true
}

func (t *Trie) HasWord(word string) bool {
	runes := []rune{}
	for _, r := range word {
		runes = append(runes, r)
	}

	return t.nodes[index(runes[0])].hasWord(runes)
}

func (n *trieNode) hasWord(word []rune) bool {
	if n == nil {
		return false
	}

	if len(word) == 1 {
		return n.terminatesHere
	}

	return n.nodes[index(word[1])].hasWord((word[1:]))
}

func index(letter rune) int {
	switch letter {
	case '\'':
		return 0
	case 'a':
		return 1
	case 'A':
		return 1
	case 'b':
		return 2
	case 'B':
		return 2
	case 'c':
		return 3
	case 'C':
		return 3
	case 'd':
		return 4
	case 'D':
		return 4
	case 'e':
		return 5
	case 'E':
		return 5
	case 'f':
		return 6
	case 'F':
		return 6
	case 'g':
		return 7
	case 'G':
		return 7
	case 'h':
		return 8
	case 'H':
		return 8
	case 'i':
		return 9
	case 'I':
		return 9
	case 'j':
		return 10
	case 'J':
		return 10
	case 'k':
		return 11
	case 'K':
		return 11
	case 'l':
		return 12
	case 'L':
		return 12
	case 'm':
		return 13
	case 'M':
		return 13
	case 'n':
		return 14
	case 'N':
		return 14
	case 'o':
		return 15
	case 'O':
		return 15
	case 'p':
		return 16
	case 'P':
		return 16
	case 'q':
		return 17
	case 'Q':
		return 17
	case 'r':
		return 18
	case 'R':
		return 18
	case 's':
		return 19
	case 'S':
		return 19
	case 't':
	case 'T':
		return 20
	case 'u':
		return 21
	case 'U':
		return 21
	case 'v':
		return 22
	case 'V':
		return 22
	case 'w':
		return 23
	case 'W':
		return 23
	case 'x':
		return 24
	case 'X':
		return 24
	case 'y':
		return 25
	case 'Y':
		return 25
	case 'z':
		return 26
	case 'Z':
		return 26
	case ' ':
		return 27
	}

	panic(fmt.Sprintf("can't find an index for a letter '%c', code '%v'\n", letter, letter))
}
