//maps.go
//program to demonstrate maps in go

package main

import "fmt"

func main() {
    elements := map[string]map[string]string{
        "H" : map[string]string {
            "name" : "Hydrogen",
            "state" : "gas",
        },
        "He" : map[string]string {
            "name" : "Helium",
            "state" : "gas",
        },
        "Li" : map[string]string {
            "name" : "Lithium",
            "state" : "solid",
        },
        "Be" : map[string]string {
            "name" : "Beryllium",
            "state" : "solid",
        },
        "B" : map[string]string {
            "name" : "Boron",
            "state" : "gas",
        },
        "C" : map[string]string {
            "name" : "Carbon",
            "state" : "solid",
        },
        "N" : map[string]string {
            "name" : "Nitrogen",
            "state" : "gas",
        },
        "O" : map[string]string {
            "name" : "Oxygen",
            "state" : "gas",
        },
        "F" : map[string]string {
            "name" : "Flourine",
            "state" : "gas",
        },
        "Ne" : map[string]string {
            "name" : "Neon",
            "state" : "gas",
        },
        "Na" : map[string]string {
            "name" : "Sodium",
            "state" : "solid",
        },
    }

    elements["Mg"] = map[string]string {
        "name": "Magnesium",
        "state" : "solid",
    }
    fmt.Println("elements -> ", elements)
    fmt.Println("len(elemnts) ->", len(elements))
    delete(elements, "Ne")
    fmt.Println("len(elemnts) ->", len(elements))

    names, ok := elements["X"]
    fmt.Println("element X -->", names, ok)

    if _,b := elements["Na"]; b {
        fmt.Println("b -->", b)
    }

    if el,ok := elements["Li"]; ok {
        fmt.Println("name ->", el["name"], "state ->", el["state"])
    }

    slc := make([]int, 3, 9)
    fmt.Println("len(slc) =", len(slc))

    x := [6]string{"a","b","c","d","e","f"}
    slc1 := x[2:5]
    fmt.Println("slc1 ->", slc1)

    x1 := []int{
        48,96,86,68,
        57,82,63,70,
        37,34,83,27,
        19,97, 9,17,
    }
    smallest_num := int(1e9)
    for _, elem  := range(x1) {
        if elem < smallest_num {
            smallest_num = elem
        }
    }
    fmt.Println("smallest num in x1 = ", smallest_num)
}
