package main

import "fmt"

func main() {

    var s = []int{1, 2, 3}

    fmt.Println(s[1:2])
    fmt.Println(s[1:3])
    fmt.Println(s[3:])
    fmt.Println(s[3:8])
}
