package main

import (
	"fmt"
	"math"
)

type Circle struct /* implements AreaFinder */ {
	Radius float32
}

func (this Circle) Area() float32 {
	return math.Pi * this.Radius * this.Radius
}

func (this Circle) Perimeter() float32 {
	return 2 * math.Pi * this.Radius
}

type Rectangle struct {
	Height float32
	Width  float32
}

func (this Rectangle) Area() float32 {
	return this.Height * this.Width
}

func (this Rectangle) Perimeter() float32 {
	return 2 * (this.Height + this.Width)
}

/* utility functions */
type AreaFinder interface {
	Area() float32
}

func PrintArea(x AreaFinder) {
	fmt.Println("Area :", x.Area())
}

type PerimeterFinder interface {
	Perimeter() float32
}

func PrintPerimeter(x PerimeterFinder) {
	fmt.Println("Perimeter :", x.Perimeter())
}

/* interface composition */
type ShapeStatsFinder interface {
	AreaFinder
	PerimeterFinder
}

func PrintShape(x ShapeStatsFinder) {
	PrintArea(x)
	PrintPerimeter(x)
}

func main() {

	c := Circle{Radius: 12}

	var af AreaFinder
	af = c
	fmt.Println(af.Area())
	// fmt.Println("Area :", c.Area())
	/*
		PrintArea(c)
		PrintPerimeter(c)
	*/
	PrintShape(c)

	r := Rectangle{Height: 10, Width: 12}
	// fmt.Println("Area :", r.Area())
	/*
		PrintArea(r)
		PrintPerimeter(r)
	*/
	PrintShape(r)

}
