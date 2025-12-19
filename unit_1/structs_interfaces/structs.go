package main

import (
	"math"
)

const pi = 3.14

// Structs
type Rectangle struct {
	Length float32
	Width  float32
}

func (r Rectangle) Area() float32 {
	return 2 * (r.Length + r.Width)
}

type Shape interface {
	Area() float32
}

type Circle struct {
	Radius float32
}

func (c Circle) Area() float32 {
	return math.Pi * c.Radius * c.Radius
}

type Triangle struct {
	Length float32
	Height float32
}

func (t Triangle) Area() float32 {
	return (t.Height * t.Length) / 2
}

func main() {

}

// Utility functions
func parimeter(rectangle Rectangle) float32 {
	return 2 * (rectangle.Length + rectangle.Width)
}
