package main

import "fmt"

func main() {

	v := 123

	ModInterface(&v)

	fmt.Println(v)
}

func ModInterface(value interface{}) {

	vv, ok := value.(*int)

	if !ok {
		panic(vv)
	}

	*vv = 456

}