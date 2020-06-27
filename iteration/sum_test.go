package iteration

import (
	"fmt"
	"reflect"
	"testing"
)

func ExampleSum() {
	numbers := []int{1, 2, 3, 4, 5}
	sum := Sum(numbers)

	fmt.Println(sum)
	// Output: 15
}

func BenchmarkSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		numbers := []int{1, 2, 3, 4, 5}
		sum := Sum(numbers)

		fmt.Println(sum)
		// Output: 15
	}
}

func TestSum(t *testing.T) {

	t.Run("collection of any number", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4}

		result := Sum(numbers)
		expected := 10

		if result != expected {
			t.Errorf("got %q want %q", result, expected)
		}
	})
}

func ExampleSumAll() {
	slice1 := []int{1, 2}
	slice2 := []int{0, 9}
	sum := SumAll(slice1, slice2)

	fmt.Println(sum)
	// Output: [3 9]
}

func BenchmarkSumAll(b *testing.B) {
	for i := 0; i < b.N; i++ {
		slice1 := []int{1, 2}
		slice2 := []int{0, 9}
		SumAll(slice1, slice2)
	}
}

func TestSumAll(t *testing.T) {
	t.Run("should sum a single slice", func(t *testing.T) {
		slice := []int{1, 1, 1}

		result := SumAll(slice)
		expected := []int{3}

		checkSums(result, expected, t)
	})

	t.Run("should sum a number of slices", func(t *testing.T) {

		slice1 := []int{1, 2}
		slice2 := []int{0, 9}

		result := SumAll(slice1, slice2)
		expected := []int{3, 9}

		checkSums(result, expected, t)
	})
}

func ExampleSumAllTails() {
	slice1 := []int{1, 2}
	slice2 := []int{0, 9}
	sum := SumAllTails(slice1, slice2)

	fmt.Println(sum)
	// Output: [2 9]
}

func BenchmarkSumAllTails(b *testing.B) {
	for i := 0; i < b.N; i++ {
		slice1 := []int{1, 2}
		slice2 := []int{0, 9}
		SumAllTails(slice1, slice2)
	}
}

func TestSumAllTails(t *testing.T) {

	t.Run("make the sums of some slices", func(t *testing.T) {
		result := SumAllTails([]int{1, 2}, []int{0, 9})
		expected := []int{2, 9}
		checkSums(result, expected, t)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		result := SumAllTails([]int{}, []int{3, 4, 5})
		expected := []int{0, 9}
		checkSums(result, expected, t)
	})
}

func checkSums(result []int, expected []int, t *testing.T) {
	t.Helper()

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("got %v want %v", result, expected)
	}
}
