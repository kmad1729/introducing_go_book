//given an array of stock prices in increasing order of time 
//find the max profit that can be made by buying and selling the stock exactly once

package main

import "fmt"

func main() {
    var stock_arr = [10]int{
        310,
        315,
        275,
        295,
        260,
        270,
        290,
        230,
        255,
        250,
    }

    max_profit := -10000000
    min_till_now := 10000000

    for i := 0; i < len(stock_arr); i++ {
        if stock_arr[i] < min_till_now{
            min_till_now = stock_arr[i]
        }
        if stock_arr[i] - min_till_now > max_profit {
            max_profit = stock_arr[i] - min_till_now
        }
    }

    fmt.Println("max profit =", max_profit)
}
