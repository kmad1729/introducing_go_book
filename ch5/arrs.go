//program to demonstrate arras in go

package main

import "fmt"

func main() {
    var x [5]int
    x[4] = 100
    fmt.Println(x);

    var scores[5]float64;

    scores[0] = 98
    scores[1] = 93
    scores[2] = 77
    scores[3] = 82
    scores[4] = 83

    var total float64;
    for idx:=0;idx < len(scores); idx++ {
        total += scores[idx];
    }

    total /= float64(len(scores))

    fmt.Printf("Average scrore = %3.2f\n", total)


    //enumerate example
    stockPrices := [10]float64 {52.2, 84.6, 45.7, 69.7, 79.1, 67.2, 3.2, 70.6, 5.8, 65.7}

    var max_stock_price = -1.0
    for i:=0;i<len(stockPrices);i++ {
        if stockPrices[i] > max_stock_price {
            max_stock_price = stockPrices[i]
        }
    }

    fmt.Println("Max stock price= ", max_stock_price)


    day_of_week := [7]string {
        "Sunday",
        "Monday",
        "Tuesday",
        "Wednesday",
        "Thursday",
        "Friday",
        "Saturday",
    }

    for i:= 0; i < len(day_of_week);i++ {
        day := day_of_week[i]

        suff := ""
        if i ==  0 {
            suff = "st"
        }else if  i == 1 {
            suff = "nd"
        }else if i == 2 {
            suff = "rd"
        }else {
            suff = "th"
        }

        fmt.Printf("%d%s of the week = %s\n", i + 1, suff, day)
    }
}
