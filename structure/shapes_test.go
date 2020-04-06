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

	areaTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, want: 72.0},
		{name: "Circle", shape: Circle{Radius: 10}, want: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Base: 12, Height: 6}, want: 36.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.want {
				t.Errorf("%#v got %g hasArea %g", tt.shape, got, tt.want)
			}
		})
	}

}
