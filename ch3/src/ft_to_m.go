//simple program to convert distance in feet to meters

package main

import "fmt"

func main() {
    var (
        feet= 0.0
        meters = 0.0
    )
    fmt.Print("Distance in feet:")
    fmt.Scanf("%f", &feet)

    meters = 0.3 * feet
    fmt.Printf("Distance in meters: %3.2f\n", meters)
}
