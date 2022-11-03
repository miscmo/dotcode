package main

import (
	"math/rand"
	"sync"
	"time"
)

func main() {
	var rspValue []int
	var wg sync.WaitGroup
	var tmux sync.Mutex

	for i := 0; i < 100; i++ {
		var v int
		rspValue = append(rspValue, v)
		rspIdx := len(rspValue) - 1

		wg.Add(1)
		go func(value int, idx int) {
			defer wg.Done()

			time.Sleep(time.Duration(rand.Intn(100)))

			tmux.Lock()
			defer tmux.Unlock()

			rspValue[idx] = value

		}(i, rspIdx)

	}

	wg.Wait()

	print(rspValue)

	for _, v := range rspValue {
		print(v)
		print(",")
	}
}