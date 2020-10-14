package numerals

import (
	"fmt"
	"testing"
)

func TestRomanNumerals(t *testing.T) {
	testinputs := []struct {
		integer int
		numeral string
	}{
		{1, "I"},
		{2, "II"},
		{3, "III"},
		{4, "IV"},
		{5, "V"},
		{6, "VI"},
		{7, "VII"},
		{8, "VIII"},
		{9, "IX"},
		{10, "X"},
		{10, "X"},
		{14, "XIV"},
		{18, "XVIII"},
		{20, "XX"},
		{39, "XXXIX"},
		{40, "XL"},
		{47, "XLVII"},
		{49, "XLIX"},
		{50, "L"},
		{100, "C"},
		{90, "XC"},
		{400, "CD"},
		{500, "D"},
		{900, "CM"},
		{1000, "M"},
		{1984, "MCMLXXXIV"},
		{3999, "MMMCMXCIX"},
		{2014, "MMXIV"},
		{1006, "MVI"},
		{798, "DCCXCVIII"},
	}

	t.Run("should return correct roman numeral for associated integer", func(t *testing.T) {
		for _, test := range testinputs {
			t.Run(fmt.Sprintf("integer value %d is converted to numeral %q", test.integer, test.numeral), func(t *testing.T) {
				result := convertToNumeral(test.integer)
				if result != test.numeral {
					t.Errorf("expected %q, got %q", test.numeral, result)
				}
			})
		}
	})

	t.Run("should return correct integer for associated roman numeral", func(t *testing.T) {
		for _, test := range testinputs {
			t.Run(fmt.Sprintf("numeral value %q is converted to integer %d", test.numeral, test.integer), func(t *testing.T) {
				result := convertToInteger(test.numeral)
				if result != test.integer {
					t.Errorf("expected %d, got %d", test.integer, result)
				}
			})
		}
	})
}
