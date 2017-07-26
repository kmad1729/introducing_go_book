//slicetest.go
// program to use slices

package main

import "fmt"

func main() {
    test_slice := make([]string, 2, 3)

    fmt.Println(test_slice)
    test_slice[0] =  "Superman"
    test_slice[1] = "Batman"
    //test_slice[2] = "Spiderman"
    fmt.Println("after populating test_slice")
    fmt.Println(test_slice)


    days_of_week := [7]string {
        "Sunday",
        "Monday",
        "Tuesday",
        "Wednesday",
        "Thursday",
        "Friday",
        "Saturday",
    }

    first_3_days := days_of_week[:3]
    fmt.Println("first_3_days = ", first_3_days)


}
