package pkg1

import "fmt"

var dx, _ = fmt.Println("pkg1.d 变量 init")

var ax, _ = fmt.Println("pkg1.a 变量 init")


func init()  {
	fmt.Println("pkg1.a first init函数")
}

func init() {
	fmt.Println("pkg1.a second init函数")
}
