/*
Sum( numbers[] ) :
SumAll( ...numbers[] ) :
Provides examples of the following:
 array
 slices - append
	mySlice := []int{1,2,3} rather than myArray := [3]int{1,2,3}
	A slice is a descriptor of an array segment. It consists of a pointer to the array, the length of the segment, and its capacity (the maximum length of the segment)
 range - lets you iterate over an array. Every time it is called it returns two values, the index and the value.
 	_ - blank identifier (can be used to ignore the index)
 variadic function - can be called with any number of trailing arguments.
 make
 len
 reflect.DeepEqual
 test coverage
*/
package iteration

func Sum(numbers []int) int {
	total := 0

	for _, number := range numbers {
		total += number
	}

	return total
}

func SumAll(numbersToSum ...[]int) []int {
	var sums []int

	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}

	return sums
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}

	return sums
}
