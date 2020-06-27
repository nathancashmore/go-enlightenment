package iteration

import (
	"fmt"
	"strings"
	"testing"
)

func ExampleRepeat() {
	fmt.Printf(Repeat("a", 5))
	// Output: aaaaa
}

func Benchmark(b *testing.B) {
	for i := 0; i < b.N; i++ {

		// Just to be clear this is about 200 ns
		Repeat("a", 5)

		// vs the library which is 60 ns !
		strings.Repeat("b", 5)
	}
}

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 3)
	expected := "aaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func TestRepeatWithModule(t *testing.T) {
	repeated := strings.Repeat("b", 6)
	expected := "bbbbbb"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}
