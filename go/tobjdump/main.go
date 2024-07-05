package main

import "unsafe"

func main() {
	n := int32(10)
	println(read(&n))

	println(size())
}

func read(p *int32) (v int32) {
	v = * p
	return
}

func size() (o uintptr) {
	o = unsafe.Sizeof(o)
	return
}



这是一个测试test
zheshige