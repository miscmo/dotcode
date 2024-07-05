package main

import "fmt"
//
//type Conn struct {
//	tag int
//}
//
//func (c *Conn) Connect(i int) {
//	c.tag = i
//	fmt.Printf("connect %v\n", c.tag)
//}
//
//func (c *Conn) Close() {
//	fmt.Printf("close %v\n", c.tag)
//}
//
func main() {

	// 1.
	//for _, v := range []int{1, 2, 3, 4, 5, 6, 7, 8} {
	//	c := &Conn{}
	//	c.Connect(v)
	//
	//	defer c.Close()
	//}

	// 2.

	a, b := 1, 2
	swap1(a, b)
	fmt.Println(a, b)
}

func swap1(a, b int) {
	a, b = b, a
}
