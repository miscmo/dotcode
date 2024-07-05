package main

import (
	"fmt"
	"sync"
)

func main() {

	datas := []int32{2, 3, 4, 5, 6, 7, 8, 1}

	var wg sync.WaitGroup
	//var tmux sync.Mutex


	for i := 0; i < 10; i++ {

		var retDatas []int32
		for j := 0; j < 10; j++ {

			wg.Add(1)

			go func() {
				defer wg.Done()

				//tmux.Lock()
				//defer tmux.Unlock()
				retDatas = append(retDatas, datas...)

			}()

		}

		wg.Wait()
		fmt.Println(retDatas)

	}

}
