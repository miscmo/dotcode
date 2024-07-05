package main

import (
	"fmt"
	_ "github.com/miscmo/dotcode/go/pkg1"
	_ "github.com/miscmo/dotcode/go/pkg2"

)

var x, _ = fmt.Println("main 变量 init")

func init()  {
	fmt.Println("main init函数")

}

func main() {
	fmt.Println("init finished")
}
