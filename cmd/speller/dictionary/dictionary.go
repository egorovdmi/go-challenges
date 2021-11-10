package dictionary

import (
	"hash/crc32"
	"strings"
)

const hashtableSize = 100000000

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

func (d *Dictionary) Load(words []string) {
	for _, w := range words {
		h := hash(w)
		if d.hashtable[h] == nil {
			d.hashtable[h] = newList(w)
			d.ListsCount++
			continue
		}

		d.hashtable[h].addToList(w)
	}

	d.Count = len(words)
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
	crc := crc32.ChecksumIEEE([]byte(strings.ToUpper(w)))
	return int(crc % hashtableSize)
}
