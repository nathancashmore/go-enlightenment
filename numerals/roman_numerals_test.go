package numerals

import "testing"

func TestRomanNumerals(t *testing.T) {

	expected := "I"
	result := convert(1)

	if result != expected {
		t.Errorf("expected %q, got %q", expected, result)
	}
}
