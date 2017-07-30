//day2.go
//more fun with functions

package main

import "fmt"


func makeEvenGenerator() func() uint {
    i := uint(0)
    return func() (r uint) {
        r = i
        i += 2
        return
    }
}


func myStepGen(step uint) func() uint {
    i := uint(0)
    return func() (r uint) {
        r = i
        i += step
        return 
    }
}


func main() {
    DELIM := "***********"
    myFunc := makeEvenGenerator()
    myFunc1 := makeEvenGenerator()

    fmt.Println(myFunc())
    fmt.Println(myFunc())
    fmt.Println(myFunc())
    fmt.Println(myFunc())
    fmt.Println(DELIM)

    fmt.Println(myFunc1())
    fmt.Println(myFunc1())
    fmt.Println(myFunc1())
    fmt.Println(DELIM)

    myFunc3 := myStepGen(15)
    for i:=0; i< 5;i++ {
        fmt.Println(i, "-->", myFunc3())
    }
    fmt.Println(DELIM)


}
