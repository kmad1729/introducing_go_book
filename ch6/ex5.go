//ex5.go
//implement a bad fib generator using recurtion

package main

import "fmt"


const (
    MAX_IND = 51
)

func fib(n uint) uint {
    if (n <= 1) {
        return n
    }
    return fib(n - 1) + fib(n - 2)
}


func main() {
    fmt.Printf("First %d fibonacci numbers -->\n", MAX_IND)
    for i:=1; i < MAX_IND; i++ {
        fmt.Println(i,"-->",fib(uint(i)))
    }
}
