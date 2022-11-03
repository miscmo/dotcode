package main

import (
	"fmt"
)

func main() {
	a := 0
	fn := func() int {
		a++
		return a
	}
	ch := make(chan int, 1)
	chh := make(chan int, 3)
	for i := 0; i < 3; i++ {
		go func(j int) {
			for {
				ch <- 1
				n := fn()
				if n > 100 {
					chh <- 1
					<-ch
					return
				}
				fmt.Println("go", j, n)
				<-ch
			}
		}(i)
	}
	for i := 0; i < 3; i++ {
		aa, ok := <-chh
		fmt.Println(aa, ok)
	}
}


