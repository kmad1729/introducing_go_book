//test_slices.go
//program to test out slices in golang

package main
import "fmt"
import "time"

const SLEEP_INTERVAL time.Duration =  .1e9
const MAXIND = 10

func main() {
    fmt.Println(time.Now())
    fmt.Println("Sleeping for", SLEEP_INTERVAL)
    time.Sleep(SLEEP_INTERVAL)
    fmt.Println(time.Now())

    time_stamp_slice := make([]time.Time, MAXIND)
    for i:= 0; i< MAXIND; i++ {
        time_stamp_slice[i] = time.Now()
        fmt.Println(i, "recorded time stamp. Sleeping for", SLEEP_INTERVAL)
        time.Sleep(SLEEP_INTERVAL)

    }

    for _, t := range(time_stamp_slice) {
        fmt.Println("time stamp =", t)
    }

    emptySlice := make([]int64,10, 12)
    fmt.Println("emptySlice =", emptySlice)
    emptySlice[9] = 42
    fmt.Println("emptySlice =", emptySlice)
    newSlice := append(emptySlice, 4,5,6,6,7,)
    fmt.Println("newSlice =", newSlice)


    slice1 := []int{1,2,3}
    slice2 := make([]int, 2)
    copy(slice1, slice2)
    fmt.Println(slice1, slice2)

}

