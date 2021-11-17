package strings_test

import (
	"testing"

	"github.com/egorovdmi/go-challenges/internals/platform/strings"
)

func TestIsPalindromePermutation(t *testing.T) {
	isPalindrome, palindromes := strings.IsPalindromePermutation("Tact Cam")

	if !isPalindrome {
		t.Errorf("IsPalindromePermutation result was incorrect, got: %v, want: %v.", isPalindrome, true)
	}

	expected := 1
	if len(palindromes) != expected {
		t.Errorf("IsPalindromePermutation result was incorrect, got: %v, want: %v.", len(palindromes), expected)
	}

	t.Logf("Generated palindrome string: %v", palindromes[0])

	isPalindrome, _ = strings.IsPalindromePermutation("Tact Coa123")

	if isPalindrome {
		t.Errorf("IsPalindromePermutation result was incorrect, got: %v, want: %v.", isPalindrome, false)
	}
}
