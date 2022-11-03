package main

import "fmt"

func main() {
    mapSliceObj := map[int][]string{}

    firstElem := mapSliceObj[1]

    firstElem = append(firstElem, "test")

    mapSliceObj[1] = firstElem

    fmt.Printf("mapSlice: %+v", mapSliceObj)

    return
}
