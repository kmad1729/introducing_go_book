package main

import (
	"fmt"
	"math"
)

type BNode struct {
	left  *BNode
	right *BNode
	data  int64
}

type Circle struct {
	radius float64
	x      float64
	y      float64
}

func insert(n **BNode, data int64) {
	if *n == nil {
		*n = new(BNode)
		(*n).data = data
		return
	} else {
		if data <= (*n).data {
			insert(&(*n).left, data)
		} else {
			insert(&(*n).right, data)
		}
	}

}

func printInOrder(n *BNode) {
	if n != nil {
		printInOrder(n.left)
		fmt.Printf("%v ", n.data)
		printInOrder(n.right)
	}
}

func (c *Circle) Area() (A float64) {
	A = math.Pi * c.radius * c.radius
	return
}

func main() {

	c1 := Circle{radius: 2, x: 0, y: 0}

	fmt.Printf("%+v\n", c1)

	fmt.Printf("Area of c1= %3.2f\n", c1.Area())

	var root *BNode = nil
	insert(&root, 42)
	insert(&root, 2)
	insert(&root, -2)
	insert(&root, 97)

	fmt.Printf("%v\n", root.data)

	printInOrder(root)
	fmt.Println()

}
