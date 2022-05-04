package main

import "fmt"

func main() {
    u := []int{11, 12, 13, 14, 15}
    s := u[1:3]
    fmt.Printf("slice(len=%d, cap=%d): %v\n", len(s), cap(s), s)    // [12, 13]

    s = append(s, 24)
    fmt.Printf("after append 24, array: %v\n", u)
    fmt.Printf("after append 24, slice(len=%d, cap=%d): %v\n", len(s), cap(s), s)

    s = append(s, []int{25, 26}...)
    fmt.Printf("after append {25, 26}, array: %v\n", u)
    fmt.Printf("after append {25, 26}, slice(len=%d, cap=%d): %v\n", len(s), cap(s), s)

    s[0] = 22
    fmt.Printf("after reassign 1st elem of slice, array: %v\n", u)
    fmt.Printf("after reassign 1st elem of slice, slice(len=%d, cap=%d): %v\n", len(s), cap(s), s)
}
