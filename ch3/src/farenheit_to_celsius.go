//program to convert temperature from Faranheit to Celsius

package main

import "fmt"

func main() {
    var (
        faranheit = 0.0
        celsius = 0.0
    )
    fmt.Print("Temperature in faranheit(F): ")
    fmt.Scanf("%f", &faranheit)
    celsius = (faranheit - 32) * (5.0/9)
    fmt.Printf("Temperature in Celsius: %3.2f\n", celsius)
}
