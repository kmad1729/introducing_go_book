//ex11.go
//write a program that swaps 2 integers


package main

import "fmt"

func swap(x *int, y *int) {
    tmp := *x
    *x = *y
    *y = tmp
}


func main() {
    var (
        x = 1
        y = 2
    )

    fmt.Printf("Before x = %d, y = %d\n",x,y)
    swap(&x, &y)
    fmt.Printf("After x = %d, y = %d\n",x,y)
}
