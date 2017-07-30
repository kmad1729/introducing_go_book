//day2_panic_recover.go
//program displaying panic and recover options

package main

import "fmt"

func main() {
    defer func() {
        recover_op := recover()
        fmt.Println("recover_op =", recover_op)
    }()
    fmt.Println("This will be run :)")
    panic("Panicing!!")

    fmt.Println("This will never be run :(")
}
