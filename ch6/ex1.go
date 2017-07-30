//sum is a function that takes a slice of numbers and adds them together
//implement sum

package main

import "fmt"

func sum(nums []int) int {
    result := 0
    for _,v := range(nums) {
        result += v
    }
    return result
}

func variadicSum(nums...int) int {
    result := 0
    for _,v := range(nums) {
        result += v
    }
    return result
}

func main() {
    myArr := [5]int {
        1,
        2,
        3,
        4,
        5,
    }

    fmt.Println("sum of myArr =", sum(myArr[:]))
    fmt.Println("variadicSum(1,2,3,4,5,6,7) =",variadicSum(1,2,3,4,5,6,7))
    fmt.Println("variadicSum(myArr[:]...) =", variadicSum(myArr[:]...))
}
