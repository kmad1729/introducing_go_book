//ex7_8_9.go
//7 -> get memory address of a varaible
//8 -> assign a value of to a pointer
//9 -> how do you create a new pointer

package main

import "fmt"

func main() {
    my_map := make(map[string]string)
    my_map["Superman"] = "Clark Kent"
    my_map["Batman"] = "Bruce Wayne"

    fmt.Println("my_map -> ", my_map)

    fmt.Printf("my_map memory loction -> %p\n", &my_map)

    var my_int_ptr *int = new(int)
    *my_int_ptr = 8
    fmt.Printf("value of my_int_ptr %d\n", *my_int_ptr)

    var my_lst_ptr *[5]int = new([5]int)
    my_lst_ptr[3] = 32

    fmt.Println("my_lst ->", *my_lst_ptr)

}
