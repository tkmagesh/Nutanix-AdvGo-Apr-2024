package main

import (
	"fmt"
	"math"
)

type Product struct {
	Id   int
	Name string
	Cost float32
}

// sprint-1
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// sprint-2
type Rectangle struct {
	Length  float64
	Breadth float64
}

func (r Rectangle) Area() float64 {
	return r.Length * r.Breadth
}

/*
func PrintArea(x interface{}) {
	switch val := x.(type) {
	case Circle:
		fmt.Println("Area :", val.Area())
	case Rectangle:
		fmt.Println("Area :", val.Area())
	default:
		fmt.Println("Incompatible argument type")
	}
}
*/

/*
func PrintArea(x interface{}) {
	switch val := x.(type) {
	case interface{ Area() float64 }:
		fmt.Println("Area :", val.Area())
	default:
		fmt.Println("Incompatible argument type")
	}
}
*/

type AreaFinder interface {
	Area() float64
}

func PrintArea(x AreaFinder) {
	fmt.Println("Area :", x.Area())
}

// Sprint-3 (introduce perimeter)
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Length + r.Breadth)
}

type PerimeterFinder interface {
	Perimeter() float64
}

func PrintPerimeter(x PerimeterFinder) {
	fmt.Println("Perimeter :", x.Perimeter())
}

/*
func PrintStats(x interface {
	AreaFinder
	PerimeterFinder
}) {
	PrintArea(x)
	PrintPerimeter(x)
}
*/

type ShapeStatsFinder interface {
	AreaFinder
	PerimeterFinder
}

func PrintStats(x ShapeStatsFinder) {
	PrintArea(x)
	PrintPerimeter(x)
}

/* func PrintStats(x interface {
	Area() float64
	Perimeter() float64
}) {
	PrintArea(x)
	PrintPerimeter(x)
} */

func main() {

	c := Circle{Radius: 12}
	// fmt.Println("Area :", c.Area())
	/*
		PrintArea(c)
		PrintPerimeter(c)
	*/
	PrintStats(c)

	r := Rectangle{Length: 10, Breadth: 14}
	// fmt.Println("Area :", r.Area())
	/*
		PrintArea(r)
		PrintPerimeter(r)
	*/
	PrintStats(r)

	/*
		pen := Product{}
		PrintArea(pen)
	*/
}
