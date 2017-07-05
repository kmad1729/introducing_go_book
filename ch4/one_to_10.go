//program to count from 1 to 5
package main

import "fmt"

var DELIM = "####################"

func main() {
    var max_num int = 5
    fmt.Println("hello world", "here are the numbers from 1 to", max_num)
    fmt.Print(`1
2
3
4
5
`)
    fmt.Println('c')

    i := 1
    for i <= max_num {
        fmt.Println(i)
        i++
    }


    fmt.Println(DELIM)
    fmt.Println("Printing FIZZ BUZZ -->")
    var max_ind = 20
    for i:=1; i<= max_ind; i++ {
        if (i % 3 == 0 && i % 5 == 0) {
            fmt.Println("FizzBuzz")
        } else if(i % 3 == 0) {
            fmt.Println("Fizz")
        } else if (i % 5 == 0) {
            fmt.Println("Buzz")
        }else {
            fmt.Println(i)
        }
    }


    max_count := 10
    fmt.Println(DELIM)
    fmt.Println("Printing out even/odd")
    for i:=1; i<=max_count; i++ {
        fmt.Print(i)
        if i%2 == 0 {
            fmt.Println(" Even")
        } else {
            fmt.Println(" Odd")
        }
    }

    fmt.Println(DELIM)
    fmt.Println("Printing months using switch stmt")
    month_id  := 0
    L:
        for ;; {
            fmt.Print("Print the month id[1-12|0 to exit]: ")
            fmt.Scanf("%d", &month_id)
            switch month_id {
            case 1:  fmt.Println("January")
            case 2:  fmt.Println("February")
            case 3:  fmt.Println("March")
            case 4:  fmt.Println("April")
            case 5:  fmt.Println("May")
            case 6:  fmt.Println("June")
            case 7:  fmt.Println("July")
            case 8:  fmt.Println("August")
            case 9:  fmt.Println("September")
            case 10:  fmt.Println("October")
            case 11:  fmt.Println("November")
            case 12:  fmt.Println("December")
            case 0: {
                fmt.Println("exiting!")
                break L
            }
            default: fmt.Println("Please print a valid month id between 1 - 12")
            }
        }

}

