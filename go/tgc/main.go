package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	//a := make([]*int, 1e9) 	// 1.8s
	a := make([]int, 1e9) 	// 300us

	for i := 0; i < 10; i++ {
		start := time.Now()
		runtime.GC()
		fmt.Printf("GC took %s\n", time.Since(start))
	}

	runtime.KeepAlive(a)
}
