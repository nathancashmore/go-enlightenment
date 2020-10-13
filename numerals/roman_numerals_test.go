package numerals

import "testing"

func TestRomanNumerals(t *testing.T) {

	t.Run("should return correct roman numeral for input", func(t *testing.T) {
		testinputs := []struct {
			decimal int
			numeral string
		}{
			{1, "I"},
			{2, "II"},
			{3, "III"},
		}

		for _, test := range testinputs {
			result := convert(test.decimal)
			if result != test.numeral {
				t.Errorf("expected %q, got %q", test.numeral, result)
			}
		}
	})
}
