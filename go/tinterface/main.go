package main

import "fmt"

func main() {

	// 1.
	//v := 123
	//
	//ModInterface(&v)
	//fmt.Println(v)


	// 2.
	var x *int = nil

	EmptyInterface(x)
}

func ModInterface(value interface{}) {

	vv, ok := value.(*int)

	if !ok {
		panic(vv)
	}

	*vv = 456

}

func EmptyInterface(x interface{}) {
	if x == nil {
		fmt.Printf("empty interface")
		return
	} else {
		fmt.Printf("non-empty interface")
	}
}