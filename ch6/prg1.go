//prg1.go
//program demonstrating functions in golang

package main

import "fmt"

func average(xs []float64) float64 { // <<-- function signature
    total := 0.0
    for _, v := range(xs) {
        total += v
    }

    return total / float64(len(xs))
}

func printInt(x int) {
    fmt.Println(x)
}

func f2() (r int) {
    r = -1234
    return
}

func my_divMod(dividend int, divisor int) (quotient int, remainder int) {
    quotient = dividend / divisor
    remainder = dividend % divisor
    return
}

func get_sum(args ...int) int {
    total := 0
    for _, v := range(args) {
        total += v
    }
    return  total
}

func main() {
    someOtherName := []float64{98,93,77,82,83}
    fmt.Println(average(someOtherName))

    printInt(123)
    printInt(f2())
    tcs := [][4]int {
        [4]int {64,5,12,4},
        [4]int {64,4,16,0},
        [4]int {64,3,21,1},
        [4]int {64,5,12,4},
    }
    for _, tc := range tcs {
        q, r := my_divMod(tc[0], tc[1])
        fmt.Printf("my_divMod(%d, %d) =\t%d %d\n", tc[0], tc[1] ,q, r)
        fmt.Printf("expected divMod(%d, %d) = %d %d\n", tc[0], tc[1] ,tc[2], tc[3])
    }
    fmt.Println("tcs ->", tcs)

    fmt.Println("adding a check for key is present in map")
    tst_map := map[string]string {
        "H" : "Hydrogen",
    }

    fmt.Println("tst_map ->", tst_map)
    if elem, ok := tst_map["H"]; ok {
        fmt.Println("elem =", elem, "ok =", ok)
    } else {
        fmt.Println("elem =", elem, "ok =", ok)
    }

    fmt.Println("sum of nums in 1,1,2,3 ->", get_sum(1,1,2,3))
    other_arr := []int{1,2,3}
    fmt.Println(get_sum(other_arr...))

    add := func (x, y int) int {
        return x + y
    }
    fmt.Println(add(1,2))


    x := 0
    nextEven := func () int {
        x += 2
        return x
    }

    fmt.Println(nextEven())
    fmt.Println(nextEven())
}
