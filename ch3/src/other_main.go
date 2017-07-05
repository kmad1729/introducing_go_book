package main

import "fmt"

var x = "Hello World!"

func main() {
    fmt.Println(x)
    f()
    const y = "Hi I am a const!"
    fmt.Println(y)
    //y = "Chaning the cons :("
    var (
        a = 5
        b = 10
        c = 15
        mystr = "Superman"
    )
    fmt.Println("a is",a)
    fmt.Println("b is",b)
    fmt.Println("c is",c)
    fmt.Println("mystr =", mystr)
    doubler_function()
}

func f() {
    fmt.Println("Hi there!, I am function f()")
    fmt.Println(x)
    fmt.Println("Bye from function f() for now! :)")
}

func doubler_function() {
    fmt.Print("Enter a number: ")
    var input float64
    fmt.Scanf("%f", &input)
    output := input * 2
    fmt.Println("Double of",input, "=", output)
}
