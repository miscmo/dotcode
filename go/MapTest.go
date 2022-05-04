package main

import "fmt"

func main() {
    demo := map[string]bool{
        "a": false,
    }

    //错误，a存在，但是返回false
    fmt.Println(demo["a"])

    //正确判断方法
    _, ok := demo["a"]
    fmt.Println(ok)
}
