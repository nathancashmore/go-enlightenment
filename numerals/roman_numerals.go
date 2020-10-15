/*
Module to provide functionality to convert Roman Numerals to their integer equivalent and back again.
*/
package numerals

import (
	"strings"
)

type RomanNumeral struct {
	Integer uint16
	Numeral string
}

type RomanNumerals []RomanNumeral

func (r RomanNumerals) ValueOf(symbol string) uint16 {
	for _, numeral := range allRomanNumerals {
		if numeral.Numeral == symbol {
			return numeral.Integer
		}
	}
	return uint16(0)
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

func ConvertToNumeral(value uint16) string {
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

func ConvertToInteger(numerals string) uint16 {
	var total = uint16(0)

	for len(numerals) > 0 {
		var value = uint16(0)
		var noOfSymbolsToRemove = 0

		if len(numerals) > 1 {
			//Check the first pair of symbols
			value = allRomanNumerals.ValueOf(numerals[0:2])
			noOfSymbolsToRemove = 2
		}

		if value == 0 {
			//Check a single symbol as no pairs found
			value = allRomanNumerals.ValueOf(numerals[0:1])
			noOfSymbolsToRemove = 1
		}

		total += value

		//Remove the symbol found from the numerals
		numerals = numerals[noOfSymbolsToRemove:]
	}

	return total
}
