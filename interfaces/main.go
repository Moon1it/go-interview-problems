package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Circle struct {
	Radius float64
}

func (this Circle) Area() float64 {
	return math.Pi * this.Radius * this.Radius
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (this Rectangle) Area() float64 {
	return this.Width * this.Height
}

func TotalArea(shapes ...Shape) float64 {
	var sum float64 = 0
	for _, shape := range shapes {
		sum += shape.Area()
	}
	return sum
}

func main() {
	c := Circle{Radius: 2}
	r := Rectangle{Width: 3, Height: 4}
	fmt.Println(TotalArea(c, r)) // ~12.57 + 12 = 24.57
}
