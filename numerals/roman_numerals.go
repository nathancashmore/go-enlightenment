package numerals

import (
	"strings"
)

type RomanNumeral struct {
	Integer int
	Numeral string
}

type RomanNumerals []RomanNumeral

func (r RomanNumerals) ValueOf(symbol string) int {
	for _, numeral := range allRomanNumerals {
		if numeral.Numeral == symbol {
			return numeral.Integer
		}
	}
	return 0
}

var allRomanNumerals = RomanNumerals{
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

func convertToNumeral(value int) string {
	var result strings.Builder
	total := value

	for total > 0 {
		for _, value := range allRomanNumerals {
			if total > value.Integer-1 {
				result.WriteString(value.Numeral)
				total -= value.Integer
				break
			}
		}
	}
	return result.String()
}

func convertToInteger(numerals string) int {
	var total = 0
	var symbol strings.Builder

	for len(numerals) > 0 {
		var value = 0
		symbol.Reset()

		if len(numerals) > 1 {
			// First find a pair and check if they return a value
			symbol.WriteString(string(numerals[0]))
			symbol.WriteString(string(numerals[1]))

			value = allRomanNumerals.ValueOf(symbol.String())
		}

		if value == 0 {
			// Pair was not found so try a single value
			symbol.Reset()
			symbol.WriteString(string(numerals[0]))

			value = allRomanNumerals.ValueOf(symbol.String())
		}

		total += value

		//Remove the symbol found from the numerals
		numerals = numerals[symbol.Len():]
	}

	return total
}
