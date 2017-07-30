//day2_recur.go
//program to test out recursion

package main

import "fmt"

func factorial(n uint) uint {
    if n <= 0 {
        return 1
    }
    return n * factorial(n - 1)
}


func main() {
    for i:=uint(0); i < 12; i++ {
        fmt.Printf("%3d! = %d\n", i, factorial(i))
    }
}
