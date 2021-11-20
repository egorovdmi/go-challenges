package strings_test

import (
	"testing"

	"github.com/egorovdmi/go-challenges/internals/platform/strings"
	"github.com/stretchr/testify/assert"
)

func TestStringCompress(t *testing.T) {
	cases := []struct {
		input    string
		expected string
	}{
		{"aabcccccaaa", "a2b1c5a3"},
		{"aabbbbbbcccccsssaasa", "a2b6c5s3a2s1a1"},
		{"abcdefghijklmnop", "abcdefghijklmnop"},
	}

	for _, c := range cases {
		result := strings.CompressString(c.input)
		assert.Equal(t, result, c.expected)
	}
}
