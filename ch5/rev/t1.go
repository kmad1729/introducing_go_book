//t1.go
// slices demo

package main

import "fmt"

func main() {
    x := make([]int, 5, 10)
    MAX_IDX := len(x)
    for i:=0;i< MAX_IDX;i++ {
        x[i] = i * 2;
    }
    x = append(x, 42, 96)
    fmt.Println("x =", x)

    slice1 := []int{1,2,3}
    slice2 := make([]int, 2)
    copy(slice2, slice1)
    fmt.Println("slice1 ->", slice1)
    fmt.Println("slice2 ->", slice2)
    
    x2 := make(map[string]int)

    x2["key"] = 10
    fmt.Println(x2)

    x3  := make(map[int]int)

    for i:= 0; i < 5; i++ {
        x3[i*2] = i * 5
    }
    fmt.Println("x3 ->", x3)

    delete(x3, 6)
    fmt.Println("x3 ->", x3)

}
