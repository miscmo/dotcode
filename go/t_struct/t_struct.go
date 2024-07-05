package main

import (
	"fmt"
	"unsafe"
)

type Foo struct {
	A int8
	B int8
	C int8
}

type Foo2 struct {
	A int32
	B int64
	C int8
}

func main() {

	f := Foo2{}

	fmt.Println(unsafe.Sizeof(f))

	fmt.Println(unsafe.Alignof(f))

	fmt.Println(unsafe.Alignof(f.A))
	fmt.Println(unsafe.Alignof(f.B))
	fmt.Println(unsafe.Alignof(f.C))
}
