//script to convert kg to lb

package main

import "fmt"


var KG_TO_LB=3.33

func main() {
    //fmt.Println("1 kg =", KG_TO_LB, " pounds")
    fmt.Print("enter the weight in kg: ")
    var (
        kgs = 0.0
        lbs = 0.0
    )
    fmt.Scanf("%f", &kgs)
    lbs = KG_TO_LB * kgs
    fmt.Printf("%3.2f kgs = %3.2f lbs\n", kgs, lbs)
}
