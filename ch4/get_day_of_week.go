//program to print out day of the week based on the index

package main

import "fmt"

func main() {
    var day_idx  = 0
    L:
        for {
            fmt.Print("Please enter the day index [1-7|0 to exit]: ")
            fmt.Scanf("%d", &day_idx)


        }
}
