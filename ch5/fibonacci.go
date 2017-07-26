//fibonacci.go
// program to store fibonacci numbers in an array

package main

import "fmt"

const MAX_IND = 35

func main() {
    var fib [MAX_IND]int

    //generating the numbers
    fib[0] = 1
    fib[1] = 1

    for i:=2; i < MAX_IND; i++ {
        fib[i] = fib[i-1] + fib[i-2]
    }

    // printing the numbers
    num_suff := ""
    for i, num := range(fib) {
        if ( i  > 9 && i <  20) {
            num_suff = "th"
        } else {
            switch (i % 10) {
            case 0: num_suff = "st"
            case 1: num_suff = "nd"
            case 2: num_suff = "rd"
            default: num_suff = "th"
            }
        }

        fmt.Printf("%2d%s fib number = %d\n", i + 1, num_suff, num)
    }
}

