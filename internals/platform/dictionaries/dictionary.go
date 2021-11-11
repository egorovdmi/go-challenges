package dictionaries

import (
	"strings"
)

const hashtableSize = 1000000
const lettersInAlphabet = 26

type item struct {
	value string
	next  *item
}

type list struct {
	item
	last *item
}

type Dictionary struct {
	hashtable  [hashtableSize]*list
	Count      int
	ListsCount int
}

func NewDictionary() *Dictionary {
	var dict Dictionary

	return &dict
}

func (d *Dictionary) Add(word string) {
	d.Count++
	h := hash(word)

	if d.hashtable[h] != nil {
		d.hashtable[h].addToList(word)
	} else {
		d.hashtable[h] = newList(word)
		d.ListsCount++
	}
}

func (d *Dictionary) Contains(word string) bool {
	h := hash(word)

	if d.hashtable[h] == nil {
		return false
	}

	for i := &d.hashtable[h].item; i != nil; i = i.next {
		if i.value == strings.ToUpper(word) {
			return true
		}
	}

	return false
}

func newList(value string) *list {
	l := list{
		item: item{
			value: strings.ToUpper(value),
			next:  nil,
		},
		last: nil,
	}

	l.last = &l.item

	return &l
}

func (l *list) addToList(value string) {
	i := item{
		value: strings.ToUpper(value),
		next:  nil,
	}

	l.last.next = &i
	l.last = &i
}

func hash(w string) int {
	bytes := []byte(strings.ToUpper(w))
	sum := 0
	for i, b := range bytes {
		x := int(b) * pow(lettersInAlphabet, i, hashtableSize)
		sum += x
	}

	return sum % hashtableSize
}

func pow(number int, pow int, mod int) int {
	r := 1
	for i := 0; i < pow; i++ {
		r = (r * number) % mod
	}

	return r
}
