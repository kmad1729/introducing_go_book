//ex10.go
//what is the value of x after running this prog
//my ans --> 2.25

package main

import "fmt"

func square(n_ptr *float64) {
    *n_ptr = (*n_ptr) * (*n_ptr)
}


func main() {
    var x float64 =  1.5
    fmt.Printf("before x= %v\n", x)
    square(&x)
    fmt.Printf("after x= %v\n", x)
}


