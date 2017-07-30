//ex3 write a function with one variadic parameter which finds greatest number

package main

import "fmt"
import "math"

func get_greatest(n...int) int{
    greatest := math.MinInt64
    for _, v := range n {
        if v > greatest {
            greatest = v
        }
    }
    return greatest
}


func main() {
    fmt.Println("get_greatest(3,1,4,1939,42,99,-12,32,42,1,0) ->",
        get_greatest(3,1,4,1939,42,99,-12,32,42,1,0) )
}
