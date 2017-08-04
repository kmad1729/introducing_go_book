package main

import "fmt"

type Person struct {
	Name string
}

type Android struct {
	Person
	Model string
}

func (p *Person) talk() {
	fmt.Println("Hi! My name is", p.Name)
}

func (a *Android) talk() {
	fmt.Println("As-ta-la-vista! from", a.Name)
}

func main() {
	a1 := Android{Person{"kashyap"}, "T3000"}
	a1.talk()
}
