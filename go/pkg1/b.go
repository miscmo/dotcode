package pkg1

import "fmt"

var bx, _ = fmt.Println("pkg1.b 变量 init")

func init()  {
	fmt.Println("pkg1.b first init函数")
}

func init() {
	fmt.Println("pkg1.b second init函数")
}