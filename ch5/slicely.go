package main

import "fmt"
import "unsafe"

const (
    MAX_IND = 10
)


func main() {
    x := make([]int64, 5, 10)
    //a := int64(123)
    x[3] = 42
    x = append(x, 2, 4)
    //fmt.Println("len of x =", len(x))
    fmt.Printf("sizeof(x) = %d\n", unsafe.Sizeof(x))
    //fmt.Printf("sizeof(a) = %d\n", unsafe.Sizeof(a))

    _table := make([]int64, MAX_IND)
    idx := 0
    for  {
        idx += 1
        if idx >= MAX_IND {
            break
        }
        _table[idx] = int64(idx * 3)
    }
    fmt.Println("_table -->", _table)
}
