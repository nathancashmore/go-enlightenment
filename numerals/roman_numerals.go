package numerals

func convert(value int) string {
	result := ""

	for i := 0; i < value; i++ {
		result = result + "I"
	}

	return result
}
