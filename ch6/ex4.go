//ex4.go
//write a function makeOddGenerator which returns a function that generates odd numbers 
//in a sequence everytime the returned function is called


package main

import "fmt"

const (
    MAX_IND = 5
)

func makeOddGenerator() func() uint {
    var i uint = 1
    return func() (result uint) {
        result = i
        i += 2
        return
    }
}

func main() {
    odd_gen := makeOddGenerator()
    for i:=0; i<MAX_IND;i++ {
        fmt.Println(i, "-->", odd_gen())
    }
    fmt.Println("making another odd generator")
    odd_gen2 := makeOddGenerator()
    for i:=0; i<MAX_IND;i++ {
        fmt.Println(i, "-->", odd_gen2())
    }
}
