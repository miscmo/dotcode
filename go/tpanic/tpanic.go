package main

import (
	"sync"
)
//
//func main() {
//	defer println("in main")
//
//	go func() {
//		defer println("in goroutine")
//
//		panic("panic")
//	}()
//
//	time.Sleep(1 * time.Second)
//}


func main() {
	//defer fmt.Println("in main1")
	//defer fmt.Println("in main2")
	//defer func() {
	//	defer func() {
	//		// panic -> gopanic
	//		// 1.13:
	//		panic("panic again and again")
	//	}()
	//	panic("panic again")
	//}()
	//
	//panic("panic once")
	//var t sync.Mutex
	//
	//tmuxPanic(t)

	pdr()
}

func tmuxPanic(t sync.Mutex) {


	t.Lock()

	panic("panic when lock")
}

func pdr() {
	defer func() {
		if err := recover(); err != nil {
			println("recover failed ", err)
		}
	}()

	panic("panic")
}