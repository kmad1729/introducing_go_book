/*allrevly.go
  implement all the programs from ch5 till maps
  1) creating array
    array iteration
    use _ in iteration
  2) creating slices
    from arrays
    use [low:high] expression
    use make function
    use append function
    use copy function
*/

package main

import "fmt"

const DELIM  =  "###################"

func main() {
    //code up binary search

    my_sorted_arr := [8]int {
        4,
        7,
        12,
        37,
        42,
        63,
        87,
        91,
    }

    tgt_list := make([]int, len(my_sorted_arr) -3)
    myop := copy(tgt_list, my_sorted_arr[:])

    fmt.Println(DELIM)
    //output of copy is the number of elements thare copied
    fmt.Println("output of copy =",myop)

    fmt.Println("my_sorted_arr ->", my_sorted_arr)
    fmt.Println("tgt_list ->\t", tgt_list)
    my_sorted_arr[3] = 14
    tgt_list = append(tgt_list, 223)
    fmt.Println("my_sorted_arr ->", my_sorted_arr)
    fmt.Println("tgt_list ->\t", tgt_list)

    fmt.Println(DELIM)
    var (
        start = 0
        end = len(my_sorted_arr) - 1
        mid = -1
    )
    fmt.Println("my_sorted_arr ->", my_sorted_arr)

    for _, tgt := range(tgt_list) {
        start = 0
        end = len(my_sorted_arr) - 1
        fmt.Printf("Looking for tgt %d\n", tgt)
        for {
            if start > end {
                fmt.Printf("tgt %d could not be found :(\n",tgt)
                break
            }
            mid = start + (end - start)/ 2
            //fmt.Printf("mid = %d and my_sorted_arr[mid] = %d\n", mid, my_sorted_arr[mid])
            if tgt == my_sorted_arr[mid] {
                fmt.Printf("found tgt %d at idx (%d)..Yay!\n", tgt, mid)
                break
            } else if(tgt < my_sorted_arr[mid]) {
                end = mid -1
            } else {
                start = mid + 1
            }
        }
    }

    fmt.Println(DELIM)
    empty_slice := []int{1}
    empty_slice = append(empty_slice, 233)
    fmt.Println("empty_slice =", empty_slice)
}
