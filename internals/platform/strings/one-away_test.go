package strings_test

import (
	"testing"

	"github.com/egorovdmi/go-challenges/internals/platform/strings"
)

func TestOneAway(t *testing.T) {

	// inserted one symbol => expected true
	result := strings.OneAway("asd", "astd")

	if !result {
		t.Errorf("ERROR: OneAway (added one symbol) didn't return expected result. Actual: %v, Expected: %v", result, true)
	}

	// removed one symbol => expected true
	result = strings.OneAway("aba", "ab")

	if !result {
		t.Errorf("ERROR: OneAway (removed one symbol) didn't return expected result. Actual: %v, Expected: %v", result, true)
	}

	// added and changed more then one symbol => expected false
	result = strings.OneAway("aba", "apbaaaa")

	if result {
		t.Errorf("ERROR: OneAway (added and changed more then one symbol) didn't return expected result. Actual: %v, Expected: %v", result, false)
	}

	// one symbol replaced => expected true
	result = strings.OneAway("aba", "aaa")

	if !result {
		t.Errorf("ERROR: OneAway (one symbol replaced) didn't return expected result. Actual: %v, Expected: %v", result, true)
	}

	// two symbols replaced => expected false
	result = strings.OneAway("aba", "bbb")

	if result {
		t.Errorf("ERROR: OneAway (two symbols replaced) didn't return expected result. Actual: %v, Expected: %v", result, false)
	}
}
