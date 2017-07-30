//day2_defer.go
//function to demo defer in go

package main

import "fmt"
import "os"

func first() {
    fmt.Println("1st")
}

func second() {
    fmt.Println("2nd")
}

func main() {
    defer first()
    second()

    //trying to read a file
    f, _ := os.Open("my_ps_ef.txt")
    defer f.Close()


    buffSize := 100
    myBuff := make([]byte, buffSize)
    n, err := f.Read(myBuff)
    fmt.Printf("n --> %d, err --> %s\n",n, err)
    fmt.Printf("%s\n", myBuff)


}
