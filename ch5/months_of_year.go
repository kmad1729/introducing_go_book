//for loops and arrays

package main

import "fmt"

func main() {
    months := [12]string {
        "January",
        "February",
        "March",
        "April",
        "may",
        "June",
        "July",
        "August",
        "September",
        "October",
        "November",
        "December",
    }

    var month string
    var suff string

    for i :=0; i< len(months); i++ {
        month = months[i]
        if i % 10 == 0 {
            suff = "st"
        } else if  i % 10 == 1 {
            suff = "nd"
        } else if i % 10 == 2 {
            suff = "rd"
        } else {
            suff = "th"
        }

        fmt.Printf("%d%s the month of year = %s\n", i+1, suff, month)

    }
}
