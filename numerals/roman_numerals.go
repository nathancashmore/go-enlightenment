package numerals

import "strings"

func convert(value int) string {

	type RomanNumeral struct {
		Decimal int
		Numeral string
	}

	var allRomanNumerals = []RomanNumeral{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	var result strings.Builder
	total := value

	for total > 0 {
		for _, value := range allRomanNumerals {
			if total > value.Decimal-1 {
				result.WriteString(value.Numeral)
				total -= value.Decimal
				break
			}
		}
	}
	return result.String()
}
