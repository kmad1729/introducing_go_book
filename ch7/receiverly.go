package main

import (
	"fmt"
	"math"
)

type Circle struct {
	x, y, r float64
}

type Rectangle struct {
	bottomLeft_x, bottomLeft_y, topRight_x, topRight_y float64
}

func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

func dist(x1, y1, x2, y2 float64) float64 {
	return math.Sqrt(((x2 - x1) * (x2 - x1)) + ((y2 - y1) * (y2 - y1)))
}

func (r *Rectangle) area() float64 {
	l := dist(r.bottomLeft_x, r.bottomLeft_y, r.bottomLeft_x, r.topRight_y)
	w := dist(r.bottomLeft_x, r.bottomLeft_y, r.topRight_x, r.bottomLeft_y)
	return l * w
}

func main() {
	c := Circle{x: 0, y: 0, r: 1}
	fmt.Printf("area of circle c = %3.2f\n", c.area())

	r := Rectangle{bottomLeft_x: 0, bottomLeft_y: 0, topRight_x: 3, topRight_y: 4}

	fmt.Printf("area of rectangle r = %3.2f\n", r.area())
}
