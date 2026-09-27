package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
	perimeter() float64
}

type Circle struct {
	radius float64
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

type Rectangle struct {
	length, width float64
}

func (r Rectangle) area() float64 {
	return r.length * r.width
}

func (r Rectangle) perimeter() float64 {
	return 2 * (r.length + r.width)
}
func main() {
	circleInstance := Circle{5}

	rectangleInstance := Rectangle{10, 15}

	fmt.Println("This is the Area of a Circle >> %.2f", circleInstance.area())
	fmt.Println("This is the Area of a Rectangle >> %.2f", rectangleInstance.area())

}
