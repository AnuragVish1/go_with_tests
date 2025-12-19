package main

import (
	"testing"
)

func TestCalcParameter(t *testing.T) {
	t.Run("Claculating parimeter of rectangle", func(t *testing.T) {
		rectangle := Rectangle{2, 5}
		got := parimeter(rectangle)
		var want float32 = 14

		if got != want {
			t.Errorf("Got %f want %f", got, want)
		}
	})
}

func TestArea(t *testing.T) {
	testArea := []struct {
		name  string
		shape Shape
		want  float32
	}{
		{name: "Rectangle", shape: Rectangle{Length: 1, Width: 2}, want: 6},
		{name: "Circle", shape: Circle{Radius: 10}, want: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Length: 2, Height: 2}, want: 2},
	}

	for _, tt := range testArea {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("%#v %v want %v ", tt.name, got, tt.want)
		}
	}

}

func checkIfCorrect(shape Shape, want float32, t testing.TB) {
	t.Helper()
	got := shape.Area()
	if got != want {
		t.Errorf("got %f want %f", got, want)
	}
}
