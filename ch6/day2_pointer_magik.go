//day2_pointer_magic.go
//program to demo pointers

package main

import "fmt"

func zero(x *int) {
    *x = 0
}

func one(x *int) {
    *x = 1
}

func main() {
    fmt.Println("attempting to change x")
    x := 23
    fmt.Println("before x =",x)
    zero(&x)
    fmt.Println("after x =",x)

    y := new(int)
    fmt.Println("before y =",*y)
    one(y)
    fmt.Println("after y =",*y)
    fmt.Println("y =",y)
}
