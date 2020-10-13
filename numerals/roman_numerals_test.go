package numerals

import "testing"

func TestRomanNumerals(t *testing.T) {

	t.Run("should return correct roman numeral for input", func(t *testing.T) {
		testinputs := []struct {
			description string
			decimal     int
			numeral     string
		}{
			{"convert 1 -> I", 1, "I"},
			{"convert 2 -> II", 2, "II"},
			{"convert 3 -> III", 3, "III"},
			{"convert 4 -> IV", 4, "IV"},
			{"convert 5 -> V", 5, "V"},
			{"convert 6 -> VI", 6, "VI"},
			{"convert 7 -> VII", 7, "VII"},
			{"convert 8 -> VIII", 8, "VIII"},
			{"convert 9 -> IX", 9, "IX"},
			{"convert 10 -> X", 10, "X"},
		}

		for _, test := range testinputs {
			t.Run(test.description, func(t *testing.T) {
				result := convert(test.decimal)
				if result != test.numeral {
					t.Errorf("expected %q, got %q", test.numeral, result)
				}
			})
		}
	})
}
