package main

import (
	"fmt"
	"math"
)

//last month (sprint-1)
type Circle struct { //implments ShapeWithArea
	Radius float32
}

func (c Circle) Area() float32 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float32 {
	return 2 * math.Pi * c.Radius
}

//last week (sprint-2)
type Rectangle struct { //implments ShapeWithArea
	Height float32
	Width  float32
}

func (r Rectangle) Area() float32 {
	return r.Height * r.Width
}

func (r Rectangle) Perimeter() float32 {
	return 2 * (r.Height + r.Width)
}

//utility functions
//sprint-3
type ShapeWithArea interface {
	Area() float32
}

func PrintArea(x ShapeWithArea) {
	fmt.Println("Area =", x.Area())
}

type ShapeWithPerimeter interface {
	Perimeter() float32
}

func PrintPerimeter(x ShapeWithPerimeter) {
	fmt.Println("Perimeter =", x.Perimeter())
}

//sprint-4
type Shape interface {
	ShapeWithArea
	ShapeWithPerimeter
}

func PrintShape(c Shape) {
	PrintArea(c)
	PrintPerimeter(c)
}

func main() {
	//sprint-1
	c := Circle{Radius: 12}
	//fmt.Println("Area =", c.Area())
	/*
		PrintArea(c)
		PrintPerimeter(c)
	*/
	PrintShape(c)

	//sprint-2
	r := Rectangle{Height: 10, Width: 12}
	//fmt.Println("Area =", r.Area())
	/*
		PrintArea(r)
		PrintPerimeter(r)
	*/
	PrintShape(r)
}
