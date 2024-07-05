package main

import (
	"github.com/fengyoulin/hookingo"
	"reflect"
	"unsafe"
)


var hno hookingo.Hook

//go:linkname newobject runtime.newobject
func newobject(typ unsafe.Pointer) unsafe.Pointer

func fno(typ unsafe.Pointer) unsafe.Pointer {
	t := reflect.TypeOf(0)

	(*(*[2]unsafe.Pointer)(unsafe.Pointer(&t)))[1] = typ	// 相当于发射了闭包对象类型


	println(t.String())

	if fn, ok := hno.Origin().(func(typ unsafe.Pointer) unsafe.Pointer); ok {
		return fn(typ)		// 调用原runtime.newobject
	}

	return nil
}

func mc(start int) func() int {
	return func() int {
		start++
		return start
	}
}

func main() {
	var err error

	hno, err = hookingo.Apply(newobject, fno)
	if err != nil {
		panic(err)
	}

	f := mc(10)

	println(f())
}
