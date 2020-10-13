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
			{"10 gets converted to X", 10, "X"},
			{"14 gets converted to XIV", 14, "XIV"},
			{"18 gets converted to XVIII", 18, "XVIII"},
			{"20 gets converted to XX", 20, "XX"},
			{"39 gets converted to XXXIX", 39, "XXXIX"},
			{"40 gets converted to XL", 40, "XL"},
			{"47 gets converted to XLVII", 47, "XLVII"},
			{"49 gets converted to XLIX", 49, "XLIX"},
			{"50 gets converted to L", 50, "L"},
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
