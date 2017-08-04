package main

import (
	"fmt"
	"math"
)

const (
	DELIM        = "********"
	FLOAT_FORMAT = "%3.1f"
)

type Circle struct {
	r float64
}

type Triangle struct {
	a, b, c float64
}

func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

func (t *Triangle) area() float64 {
	semi_perimiter := (t.a + t.b + t.c) / 2
	prod := semi_perimiter * (semi_perimiter - t.a) * (semi_perimiter - t.b) * (semi_perimiter - t.c)
	return math.Sqrt(prod)
}

type Rectangle struct {
	length, breadth float64
}

func (r *Rectangle) area() float64 {
	return r.length * r.breadth
}

type Shape interface {
	area() float64
}

func totalArea(shapes ...Shape) (result float64) {
	result = 0.0
	for _, s := range shapes {
		result += s.area()
	}
	return
}

//using interfaces as a field
type MultiShape struct {
	shapes []Shape
}

func (ms *MultiShape) area() (result float64) {
	for _, s := range ms.shapes {
		result += s.area()
	}
	return
}

func main() {
	t1 := Triangle{a: 3, b: 4, c: 5}
	fmt.Printf(fmt.Sprintf("area of triangle 3,4,5 = %s\n", FLOAT_FORMAT), t1.area())
	//fmt.Printf(fmt.Sprintf("expected area = %s\n", FLOAT_FORMAT), 0.5*t1.a*t1.b)

	c1 := Circle{r: 1}
	fmt.Printf(fmt.Sprintf("area of circle with unit radius = %s\n", FLOAT_FORMAT), c1.area())

	r1 := Rectangle{length: 9, breadth: 2}
	fmt.Printf(fmt.Sprintf("area of rectangle 9 * 2 = %s\n", FLOAT_FORMAT), r1.area())

	fmt.Printf(fmt.Sprintf("total area of shapes c1 r1 t1 = %s\n", FLOAT_FORMAT), totalArea(&r1, &c1, &t1))

	fmt.Println(DELIM)
	ms := MultiShape{
		shapes: []Shape{
			&Circle{1},
			&Rectangle{5, 2},
			&Triangle{5, 12, 13},
		},
	}

	fmt.Println("individual areas -->")
	for i, s := range ms.shapes {
		fmt.Printf("area of shape %d = %3.2f\n", i, s.area())

	}
	fmt.Printf(fmt.Sprintf("area of multishape = %s\n", FLOAT_FORMAT), ms.area())
	fmt.Println(DELIM)

}
