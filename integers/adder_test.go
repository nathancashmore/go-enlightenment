package integers

import (
	"fmt"
	"testing"
)

func ExampleAdd() {
	sum := Add(2, 2)
	fmt.Println(sum)
	// Output: 4
}

func Benchmark(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(2, 2)
	}
}

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}
