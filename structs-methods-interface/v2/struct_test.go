package main

import (
	"testing"
)

func TestArea(t *testing.T) {
	areaTest := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 12.0, Height: 7.0}, hasArea: 84.0},
		{name: "Circle", shape: Circle{Radius: 6.0}, hasArea: 113.09733552923255},
		{name: "Triangle", shape: Triangle{Width: 12.0, Height: 8.0}, hasArea: 42},
	}

	for _, tt := range areaTest {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.hasArea {
				t.Errorf("got %g want %g", got, tt.hasArea)
			}
		})
	}
}
