//structly.go
//program demonstrating structs in golang

package main

import (
	"fmt"
	"math"
)

type Circle struct {
	x, y, r float64
}

type Rectangle struct {
	x1, y1, x2, y2 float64
}

func circleArea(c *Circle) float64 {
	return math.Pi * c.r * c.r
}

func rectangleArea(r *Rectangle) float64 {
	bottomLeft_x := r.x1
	bottomLeft_y := r.y1
	topRight_x := r.x2
	topRight_y := r.y2
	return distance(bottomLeft_x, bottomLeft_y, bottomLeft_x, topRight_y) *
		distance(bottomLeft_x, bottomLeft_y, topRight_x, bottomLeft_y)
}

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}

const (
	DELIM = "**********"
)

func main() {
	unitCircle := &Circle{x: 0, y: 0, r: 1}
	fmt.Printf("Area of unit circle = %3.3f\n", circleArea(unitCircle))

	fmt.Println(DELIM)
	smallRectangle := Rectangle{x1: 0, y1: 0, x2: 3, y2: 4}
	fmt.Printf("Area of the smallRectangle =\t%3.2f\n", rectangleArea(&smallRectangle))
	fmt.Println("expected area is\t\t12.00")
}
