package main

import (
	"fmt"
	"math"
)

// shape in the interface
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Rectangle struct
type Rectangle struct {
	Width, Height float64
}

// Rectangle's area
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Rectangle's perimeter
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// circle struct
type Circle struct {
	Radius float64
}

// circle's area
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// Triangle	Area = 1 2 bh	b = base h = height

type Triangle struct {
	Base   int
	Height int
}

func (t Triangle) Area() int {
	return 12 * t.Base * t.Height
}



func main() {
	rectangle := Rectangle{Width: 10, Height: 5}
	circle := Circle{Radius: 7}
	triangle := Triangle{Base: 4, Height: 5}
	fmt.Println("Area of Rectangle:", rectangle.Area())
	fmt.Println("Area of Circle:", circle.Area())
	fmt.Println("Area of Triangle:", triangle.Area())
}
