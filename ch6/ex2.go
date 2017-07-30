//ex2.go
//write a function that takes an integer and halves it 
//and retrun true if it was even or false if it was odd
//e.g. half(1) --> (0, false)
//     half(2) --> (1, true)

package main

import "fmt"

func half(n int) (int, bool) {
    return n / 2, n % 2 == 0
}

func main() {
    tcs := []int {
        0,
        1,
        2,
        3,
        4,
    }

    var q int
    var is_even bool
    for _,tc := range tcs {
        q, is_even = half(tc)
        fmt.Printf("half(%d) = (%d, %t)\n", tc, q, is_even)
    }
}
