package numerals

import "strings"

func convert(value int) string {
	var result strings.Builder
	total := value

	for total > 0 {
		switch {
		case total > 9:
			result.WriteString("X")
			total -= 10
		case total > 8:
			result.WriteString("IX")
			total -= 9
		case total > 4:
			result.WriteString("V")
			total -= 5
		case total > 3:
			result.WriteString("IV")
			total -= 4
		default:
			result.WriteString("I")
			total--
		}
	}

	return result.String()
}
