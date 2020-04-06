package structure

import (
	"fmt"
	"testing"
)

func TestPerimeter(t *testing.T) {

	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		result := rectangle.Perimeter()
		expected := 40.0

		if result != expected {
			t.Errorf("got %.2f want %.2f", result, expected)
		}
	})
}

func ExampleCircle_Area() {
	circle := Circle{10}
	fmt.Printf("%v", circle.Area())
	// Output: 314.1592653589793
}

func BenchmarkCircle_Area(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Circle{10}.Area()
	}
}

func TestArea(t *testing.T) {

	checkArea := func(t *testing.T, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	}

	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{7.0, 9.0}
		expected := 63.0

		checkArea(t, rectangle, expected)
	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{10}
		expected := 314.1592653589793

		checkArea(t, circle, expected)
	})
}
