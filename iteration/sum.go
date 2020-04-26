/*
Sum( numbers[] ) :
SumAll( ...numbers[] ) :
Provides examples of the following:
 array
 slices - append
	mySlice := []int{1,2,3} rather than myArray := [3]int{1,2,3}
 range
 _ - blank identifier
 variadic function
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
