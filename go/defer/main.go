package main

import "fmt"

type Conn struct {
	tag int
}

func (c *Conn) Connect(i int) {
	c.tag = i
	fmt.Printf("connect %v\n", c.tag)
}

func (c *Conn) Close() {
	fmt.Printf("close %v\n", c.tag)
}

func main() {

	for _, v := range []int{1, 2, 3, 4, 5, 6, 7, 8} {
		c := &Conn{}
		c.Connect(v)

		defer c.Close()
	}
}
