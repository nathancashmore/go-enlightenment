package numerals

func convert(value int) string {
	result := ""
	total := value

	switch value {
	case 4:
		return "IV"
	case 9:
		return "IX"
	default:
		if total >= 5 {
			result = result + "V"
			total = total - 5
		}

		for i := 0; i < total; i++ {
			result = result + "I"
		}
	}

	return result
}
