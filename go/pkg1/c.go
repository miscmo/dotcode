package pkg1

import "fmt"

var cx, _ = fmt.Println("pkg1.c 变量 init")

func init()  {
	fmt.Println("pkg1.c first init函数")
}

func init() {
	fmt.Println("pkg1.c second init函数")
}