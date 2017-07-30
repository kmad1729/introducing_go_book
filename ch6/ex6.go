//ex6.go
//what are defer, panic and recover
//how to you recover from a runtime panic
//Answer --> To recover from a runtime panic, we start off main with a deferred function 

package main

import "fmt"

func main() {
    defer func() {
        recovery_output := recover()
        fmt.Println("recovery_output -->", recovery_output)
    }()


    panic("NOOOO!")
    var my_arr [4]int
    my_arr[2] =  5
    my_slc := my_arr[:]
    my_slc = append(my_slc, 5 ,4,5)
    fmt.Println("my_arr -->", my_arr)
    fmt.Println("my_arr -->", my_slc)
}
