package injection

import (
	"bytes"
	"fmt"
	"testing"
)

func ExampleGreet() {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Chris")

	fmt.Printf(buffer.String())
	// Output: Hello, Chris
}

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}

	Greet(&buffer, "Chris")

	got := buffer.String()
	want := "Hello, Chris"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
