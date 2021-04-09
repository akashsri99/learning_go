package interfaces

import "math"

type Shape interface {
	Area() float64
}

type Circle struct {
	Radius float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (rectangle Rectangle) Area() float64 {
	return rectangle.Height * rectangle.Width
}

func (circle Circle) Area() float64 {
	return math.Pi * circle.Radius * circle.Radius
}
