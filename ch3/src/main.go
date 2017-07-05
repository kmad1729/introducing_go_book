package main

import "fmt"

func main() {
    var x string = "Hello World"
    fmt.Println(x)
    var y string
    y = "Bye World!"
    fmt.Println(y)

    var x1 string
    x1 = "This"
    x1 = x1 + " string"
    x1 = x1 + " is build in multiple lines"
    fmt.Println(x1)

    fmt.Println("trying", "string", "concatination 1st way")
    fmt.Println("trying" + "string concatination 2nd way (" +  x1 + ")")

    fmt.Println("(" + x + ") == " + "(" + y  + ")? -->" , x==y)
    pi := 3.141
    fmt.Println("Pi =", pi)

    dogsName := "Tommy"
    fmt.Println("My dog's name is", dogsName)

}
