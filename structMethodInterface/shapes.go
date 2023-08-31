package shapes

import "math"

// A little note - this is a method structure. it needs the reciverName to tie it
// to its struct as a method. Thats how it differs from a regular function
//
// func (receiverName ReceiverType) MethodName(args)
//
//Example:
//
// func (r Rectangle) Area() float64 {
// 	return 0
// }

// Shape interface
type Shape interface {
	Area() float64
}

// Rectangle declaration and methods
type Rectangle struct {
	Width  float64
	Height float64
}

func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Circle declaration and methods
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pow(c.Radius, 2) * math.Pi
}

// Triangle declaration and methods
type Triangle struct {
	Base   float64
	Height float64
}

func (t Triangle) Area() float64 {
	return 0.5 * t.Base * t.Height
}
