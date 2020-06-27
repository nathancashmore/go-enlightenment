package structure

import (
	"fmt"
	"testing"
)

func ExampleCircle_Perimeter() {
	fmt.Printf("%v", Circle{10}.Perimeter())
	// Output: 62.83185307179586
}

func ExampleRectangle_Perimeter() {
	fmt.Printf("%v", Rectangle{10, 10}.Perimeter())
	// Output: 40
}

func ExampleTriangle_Perimeter() {
	fmt.Printf("%v", Triangle{10, 10, 10, 10}.Perimeter())
	// Output: 30
}

func Benchmark(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Circle{10}.Perimeter()
	}
}

func TestPerimeter(t *testing.T) {
	tests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, want: 36.0},
		{name: "Circle", shape: Circle{Radius: 10}, want: 62.83185307179586},
		{name: "Triangle", shape: Triangle{Base: 12, SideA: 6, SideB: 6}, want: 24.0},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := test.shape.Perimeter()
			if got != test.want {
				t.Errorf("%#v got %g hasPerimeter %g", test.shape, got, test.want)
			}
		})
	}
}

func ExampleCircle_Area() {
	fmt.Printf("%v", Circle{10}.Area())
	// Output: 314.1592653589793
}

func ExampleRectangle_Area() {
	fmt.Printf("%v", Rectangle{10, 10}.Area())
	// Output: 100
}

func ExampleTriangle_Area() {
	fmt.Printf("%v", Triangle{Base: 10, Height: 10}.Area())
	// Output: 50
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
