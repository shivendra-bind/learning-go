package perimeter

import "math"

type Ractangle struct {
	width  float64
	height float64
}

type Shape interface {
	Area() float64
}

func (r Ractangle) Area() float64 {
	return r.height * r.width
}

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

type Triangle struct {
	Base   float64
	Height float64
}

func (t Triangle) Area() float64 {
	return (t.Base * t.Height) / 2
}
