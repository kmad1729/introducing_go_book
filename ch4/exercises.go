//print all the numbers evenly divisible by 3

package main

import "fmt"

func main() {
    for i:=1; i < 101; i++ {
        if i %  3 == 0 {
            fmt.Printf("%2d is a mulple of 3\n", i)
        }
    }
}
